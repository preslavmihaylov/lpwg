package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell"
)

// Extra Challenges:
// 1. Prettify the game frame
// 2. Print an error message if the screen size is not big enough
// 3. Print the score in real-time somewhere outside the game frame
// 4. Change the game colors. E.g. make the snake green and the apple red
// 5. Make the apple "timed". It should regenerate itself after a given timeout
// 6. Spawn apples and occassionally spawn "pears" which provide a higher score when eaten
// 7*. The terminal cell has higher height than width.
//     Make the snake parts "bigger" by taking up two adjacent cells instead of one
// 8*. Make a variant of the game where the snake reappears at the other end
//     of the screen if it hits the wall
// 9*. Make the game dynamically adjust when the screen is resized

// This symbol looks nicer, but doesn't render well on some terminals.
// const SnakeSymbol = 0x2588
const SnakeSymbol = '*'
const AppleSymbol = 0x25CF
const GameFrameWidth = 30
const GameFrameHeight = 15
const GameFrameSymbol = '║'

type Point struct {
	row, col int
}

type Snake struct {
	parts          []*Point
	velRow, velCol int
	symbol         rune
}

type Apple struct {
	point  *Point
	symbol rune
}

var screen tcell.Screen
var snake *Snake
var apple *Apple
var score int
var pointsToClear []*Point
var isGameOver bool
var isGamePaused bool
var debugLog string

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

	screenWidth, screenHeight := screen.Size()
	PrintStringCentered(screenHeight/2, screenWidth/2, "Game Over!")
	PrintStringCentered(screenHeight/2+1, screenWidth/2,
		fmt.Sprintf("Your score is %d...", score))
	screen.Show()
	time.Sleep(3 * time.Second)
	screen.Fini()
}

func UpdateState() {
	if isGamePaused {
		return
	}

	UpdateSnake()
	UpdateApple()
}

func UpdateSnake() {
	head := GetSnakeHead()
	snake.parts = append(snake.parts, &Point{
		row: head.row + snake.velRow,
		col: head.col + snake.velCol,
	})

	if !AppleIsInsideSnake() {
		snake.parts = snake.parts[1:]
	} else {
		score++
	}

	if IsSnakeHittingWall() || IsSnakeEatingItself() {
		isGameOver = true
	}
}

func IsSnakeHittingWall() bool {
	head := GetSnakeHead()
	return head.row < 0 ||
		head.row >= GameFrameHeight ||
		head.col < 0 ||
		head.col >= GameFrameWidth
}

func IsSnakeEatingItself() bool {
	head := GetSnakeHead()
	for _, p := range snake.parts[:GetSnakeHeadIndex()] {
		if p.row == head.row && p.col == head.col {
			return true
		}
	}

	return false
}

func GetSnakeHeadIndex() int {
	return len(snake.parts) - 1
}

func GetSnakeHead() *Point {
	return snake.parts[len(snake.parts)-1]
}

func UpdateApple() {
	for AppleIsInsideSnake() {
		apple.point.row, apple.point.col =
			rand.Intn(GameFrameHeight), rand.Intn(GameFrameWidth)
	}
}

func AppleIsInsideSnake() bool {
	for _, p := range snake.parts {
		if p.row == apple.point.row && p.col == apple.point.col {
			return true
		}
	}

	return false
}

func DrawState() {
	if isGamePaused {
		return
	}

	ClearScreen()
	PrintString(0, 0, debugLog)
	PrintGameFrame()

	DrawSnake()
	DrawApple()

	PrintFilledRect(20, 20, 1, 1, ' ')
	screen.Show()
}

func ClearScreen() {
	for _, p := range pointsToClear {
		PrintFilledRectInGameFrame(p.row, p.col, 1, 1, ' ')
	}

	pointsToClear = []*Point{}
}

func DrawSnake() {
	for _, p := range snake.parts {
		PrintFilledRectInGameFrame(p.row, p.col, 1, 1, snake.symbol)
		pointsToClear = append(pointsToClear, p)
	}
}

func DrawApple() {
	PrintFilledRectInGameFrame(apple.point.row, apple.point.col, 1, 1, apple.symbol)
	pointsToClear = append(pointsToClear, apple.point)
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
	snake = &Snake{
		parts: []*Point{
			{row: 9, col: 3}, // tail
			{row: 8, col: 3},
			{row: 7, col: 3},
			{row: 6, col: 3},
			{row: 5, col: 3}, // head
		},
		velRow: -1,
		velCol: 0,
		symbol: SnakeSymbol,
	}

	apple = &Apple{
		point:  &Point{row: 10, col: 10},
		symbol: AppleSymbol,
	}
}

func HandleUserInput(key string) {
	if key == "Rune[q]" {
		screen.Fini()
		os.Exit(0)
	} else if key == "Rune[p]" {
		isGamePaused = true
	} else if key == "Rune[w]" && snake.velRow != 1 {
		snake.velRow = -1
		snake.velCol = 0
	} else if key == "Rune[a]" && snake.velCol != 1 {
		snake.velRow = 0
		snake.velCol = -1
	} else if key == "Rune[s]" && snake.velRow != -1 {
		snake.velRow = 1
		snake.velCol = 0
	} else if key == "Rune[d]" && snake.velCol != -1 {
		snake.velRow = 0
		snake.velCol = 1
	}
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
	gameFrameTopLeftRow, gameFrameTopLeftCol := GetGameFrameTopLeft()
	row, col := gameFrameTopLeftRow-1, gameFrameTopLeftCol-1
	width, height := GameFrameWidth+2, GameFrameHeight+2

	PrintUnfilledRect(row, col, width, height, GameFrameSymbol)
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
	r, c := GetGameFrameTopLeft()
	PrintFilledRect(row+r, col+c, width, height, ch)
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

	for r := 1; r < height-1; r++ {
		screen.SetContent(col, row+r, ch, nil, tcell.StyleDefault)
		screen.SetContent(col+width-1, row+r, ch, nil, tcell.StyleDefault)
	}

	for c := 0; c < width; c++ {
		screen.SetContent(col+c, row+height-1, ch, nil, tcell.StyleDefault)
	}
}

func GetGameFrameTopLeft() (int, int) {
	screenWidth, screenHeight := screen.Size()
	return screenHeight/2 - GameFrameHeight/2, screenWidth/2 - GameFrameWidth/2
}
