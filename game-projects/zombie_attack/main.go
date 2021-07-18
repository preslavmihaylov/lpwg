package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell"
)

// Extra Challenges:
// * Prettify the game frame
// * Optimize clearing the screen
// * Colors
// * "Wall health"
// * Player health
// * Ammo & Reloading
// * Bonuses
// * Tougher zombies

const GameFrameWidth = 80
const GameFrameHeight = 25
const GameFrameSymbol = '║'

type Point struct {
	row, col int
	symbol   rune
}

type GameObject struct {
	points         []*Point
	velRow, velCol int
}

var screen tcell.Screen
var isGamePaused bool
var isGameOver bool
var debugLog string
var score int

var player *GameObject
var zombies []*GameObject
var bullets []*GameObject

func main() {
	rand.Seed(time.Now().UnixNano())

	InitScreen()
	InitGameState()
	inputChan := InitUserInput()

	for !isGameOver {
		HandleUserInput(ReadInput(inputChan))
		UpdateState()
		DrawState()

		time.Sleep(75 * time.Millisecond)
	}

	width, height := screen.Size()
	PrintStringCentered(height/2, width/2, "Game over!")
	PrintStringCentered(height/2+1, width/2, fmt.Sprintf("Your score is %d...", score))
	screen.Show()

	time.Sleep(3 * time.Second)
	screen.Fini()
}

func UpdateState() {
	if isGamePaused {
		return
	}

	MoveGameObjects(append(append(zombies, bullets...), player))
	UpdateZombies()
	CollisionDetection()
}

func UpdateZombies() {
	spawnChance := rand.Intn(100)
	if spawnChance < 5 {
		SpawnZombie()
	}
}

func SpawnZombie() {
	originRow, originCol := rand.Intn(GameFrameHeight-3), GameFrameWidth-2
	zombies = append(zombies, &GameObject{
		points: []*Point{
			{row: originRow, col: originCol, symbol: '0'},
			{row: originRow + 1, col: originCol, symbol: '|'},
			{row: originRow + 1, col: originCol - 1, symbol: '\\'},
			{row: originRow + 2, col: originCol, symbol: '|'},
			{row: originRow + 3, col: originCol - 1, symbol: '/'},
			{row: originRow + 3, col: originCol + 1, symbol: '\\'},
		},
		velRow: 0, velCol: -1,
	})
}

func MoveGameObjects(objs []*GameObject) {
	for _, obj := range objs {
		for i := range obj.points {
			obj.points[i].row += obj.velRow
			obj.points[i].col += obj.velCol
		}
	}
}

func CollisionDetection() {
	RemoveObjectsOutOfBounds()
	HandleZombiePlayerCollision()
	HandleZombieBulletCollision()

}

func RemoveObjectsOutOfBounds() {
	ObjectOutOfBoundsCollision(zombies, true, func(idx int) {
		isGameOver = true
	})

	bulletsToRemove := []*GameObject{}
	ObjectOutOfBoundsCollision(bullets, false, func(idx int) {
		bullets = append(bullets[:idx], bullets[idx+1:]...)
	})

	bullets = RemoveGameObjects(bullets, bulletsToRemove)
}

func HandleZombiePlayerCollision() {
	for _, z := range zombies {
		if AreObjectsColliding(player, z, 1) {
			isGameOver = true
		}
	}
}

func HandleZombieBulletCollision() {
	bulletsToRemove := []*GameObject{}
	zombiesToRemove := []*GameObject{}
	for _, b := range bullets {
		for _, z := range zombies {
			if AreObjectsColliding(b, z, 1) {
				bulletsToRemove = append(bulletsToRemove, b)
				zombiesToRemove = append(zombiesToRemove, z)
				score++
				break
			}
		}
	}

	bullets = RemoveGameObjects(bullets, bulletsToRemove)
	zombies = RemoveGameObjects(zombies, zombiesToRemove)
}

func RemoveGameObjects(source, toRemove []*GameObject) []*GameObject {
	result := []*GameObject{}
	for _, obj1 := range source {
		removed := false
		for _, obj2 := range toRemove {
			if obj1 == obj2 {
				removed = true
			}
		}

		if !removed {
			result = append(result, obj1)
		}
	}

	return result
}

func ObjectOutOfBoundsCollision(objs []*GameObject, lookAhead bool, callback func(int)) {
	for i, obj := range objs {
		velRow, velCol := obj.velRow, obj.velCol
		if !lookAhead {
			velRow, velCol = 0, 0
		}
		if IsObjectOutOfBounds(obj, velRow, velCol) {
			callback(i)
		}
	}
}

func AreObjectsColliding(obj1, obj2 *GameObject, radius int) bool {
	for _, p1 := range obj1.points {
		for _, p2 := range obj2.points {
			if p1.row == p2.row && math.Abs(float64(p1.col-p2.col)) <= float64(radius) {
				return true
			}
		}
	}

	return false
}

func DrawState() {
	if isGamePaused {
		return
	}

	screen.Clear()
	PrintString(0, 0, debugLog)
	PrintGameFrame()
	PrintGameObjects(append(append(zombies, bullets...), player))

	screen.Show()
}

func PrintGameObjects(objs []*GameObject) {
	for _, obj := range objs {
		for _, p := range obj.points {
			PrintFilledRectInGameFrame(p.row, p.col, 1, 1, p.symbol)
		}
	}
}

func InitScreen() {
	var err error
	screen, err = tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if err := screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	screen.SetStyle(defStyle)
}

func InitGameState() {
	player = &GameObject{
		points: []*Point{
			{row: 5, col: 1, symbol: '0'},
			{row: 6, col: 1, symbol: '|'},
			{row: 7, col: 1, symbol: '|'},
			{row: 6, col: 2, symbol: '-'},
			{row: 6, col: 3, symbol: '-'},
			{row: 6, col: 4, symbol: '-'},
			{row: 7, col: 2, symbol: '/'},
			{row: 8, col: 0, symbol: '/'},
			{row: 8, col: 2, symbol: '\\'},
		},
	}
}

func HandleUserInput(key string) {
	if key == "Rune[q]" {
		screen.Fini()
		os.Exit(0)
	} else if key == "Rune[p]" {
		isGamePaused = !isGamePaused
	} else if key == "Enter" {
		SpawnBullet(player.points[0].row+1, player.points[0].col+3)
	} else if key == "Rune[w]" && !IsObjectOutOfBounds(player, -1, 0) {
		MovePlayer(-1, 0)
	} else if key == "Rune[s]" && !IsObjectOutOfBounds(player, 1, 0) {
		MovePlayer(1, 0)
	} else if key == "Rune[a]" && !IsObjectOutOfBounds(player, 0, -1) {
		MovePlayer(0, -1)
	} else if key == "Rune[d]" && !IsObjectOutOfBounds(player, 0, 1) {
		MovePlayer(0, 1)
	}
}

func MovePlayer(velRow, velCol int) {
	for i := range player.points {
		player.points[i].row += velRow
		player.points[i].col += velCol
	}
}

func SpawnBullet(row, col int) {
	bullets = append(bullets, &GameObject{
		points: []*Point{
			{row: row, col: col, symbol: '*'},
		},
		velRow: 0, velCol: 2,
	})
}

func IsObjectOutOfBounds(obj *GameObject, velRow, velCol int) bool {
	for _, p := range obj.points {
		targetRow, targetCol := p.row+velRow, p.col+velCol
		if targetRow < 0 || targetRow >= GameFrameHeight ||
			targetCol < 0 || targetCol >= GameFrameWidth {
			return true
		}
	}

	return false
}

func InitUserInput() chan string {
	inputChan := make(chan string)
	go func() {
		for {
			switch ev := screen.PollEvent().(type) {
			case *tcell.EventKey:
				inputChan <- ev.Name()
			}
		}
	}()

	return inputChan
}

func ReadInput(inputChan chan string) string {
	var key string
	select {
	case key = <-inputChan:
	default:
		key = ""
	}

	return key
}

func PrintGameFrame() {
	frameStartRow, frameStartCol := GetFrameTopLeft()
	PrintUnfilledRect(
		frameStartRow-1, frameStartCol-1, GameFrameWidth+2, GameFrameHeight+2, GameFrameSymbol)
}

func GetFrameTopLeft() (int, int) {
	screenWidth, screenHeight := screen.Size()
	return screenHeight/2 - GameFrameHeight/2, screenWidth/2 - GameFrameWidth/2
}

func PrintStringCentered(row, col int, str string) {
	col = col - len(str)/2
	PrintString(row, col, str)
}

func PrintString(row, col int, str string) {
	for _, c := range str {
		PrintFilledRect(row, col, 1, 1, c)
		col += 1
	}
}

func PrintFilledRectInGameFrame(row, col, width, height int, ch rune) {
	frameRow, frameCol := GetFrameTopLeft()
	PrintFilledRect(frameRow+row, frameCol+col, width, height, ch)
}

func PrintFilledRect(row, col, width, height int, ch rune) {
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			screen.SetContent(col+c, row+r, ch, nil, tcell.StyleDefault)
		}
	}
}

func PrintUnfilledRect(row, col, width, height int, ch rune) {
	for c := 0; c < width; c++ {
		screen.SetContent(col+c, row, ch, nil, tcell.StyleDefault)
	}

	for r := 0; r < height; r++ {
		screen.SetContent(col, row+r, ch, nil, tcell.StyleDefault)
		screen.SetContent(col+width-1, row+r, ch, nil, tcell.StyleDefault)
	}

	for c := 0; c < width; c++ {
		screen.SetContent(col+c, row+height-1, ch, nil, tcell.StyleDefault)
	}
}
