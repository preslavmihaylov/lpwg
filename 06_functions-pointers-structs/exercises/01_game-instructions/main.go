package main

import "fmt"

func main() {
	printGameInstructions()
}

func printGameInstructions() {
	fmt.Println(`
1. Have the player select a letter of the alphabet.
2. If the letter is contained in the word/phrase, 
   all the the letters equal to it are revealed.
3. If the letter is not contained in the word/phrase, a portion of the hangman is added.
4. The game continues until:
  1. the word/phrase is guessed (all letters are revealed) – WINNER or,
  2. all the parts of the hangman are displayed – LOSER`)
}
