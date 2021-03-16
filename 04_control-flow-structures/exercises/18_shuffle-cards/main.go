package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	unshuffledDeck := strings.Split(strings.TrimSpace(line), " ")
	shuffledDeck := []string{}

	rand.Seed(time.Now().UTC().UnixNano())
	for len(unshuffledDeck) > 0 {
		cardIndex := rand.Intn(len(unshuffledDeck))
		card := unshuffledDeck[cardIndex]
		unshuffledDeck = remove(unshuffledDeck, cardIndex)
		shuffledDeck = append(shuffledDeck, card)
	}

	fmt.Println(strings.Join(shuffledDeck, " "))
}

func remove(slice []string, i int) []string {
	return append(slice[0:i], slice[i+1:]...)
}
