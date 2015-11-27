package print

import (
	"fmt"
	"time"
)

func WelcomeMessage() {
	fmt.Println("Welcome to Mastermind!")
	WhatsNext()
	prompt()
}

func QuitMessage() {
	for i := 0; i < 30; i++ {
		fmt.Println("byeeee")
		time.Sleep(30 * time.Millisecond)
	}
}

func Instructions() {
	fmt.Println("\nThe game will generate a sequence with four elements made up of: ")
	fmt.Println("(r)ed, (g)reen, (b)lue, and (y)ellow. It will look like 'rybg'.")
	fmt.Println("Your goal is to guess the code in the fewest guesses possible.")
	fmt.Println("\nGuesses must be four characters long and only contain 'r', 'g', 'b' and 'y'.")
	fmt.Println("It doesn't matter if the guess is upper case or lower case.")
	fmt.Println("The game will then tell you correct positions and correct colors.")
	fmt.Println("Good luck!\n")
	WhatsNext()
	prompt()
}

func UnknownCommand() {
	fmt.Println("Unknown command! Please try again ...")
	prompt()
}

func WhatsNext() {
	fmt.Println("Would you like to (p)lay read the (i)instructions or (q)uit?")
}

func GameStart() {
	fmt.Println("I have generated a random code four characters in length using the letters:")
	fmt.Println("'r', 'g', 'b' and 'y'")
	fmt.Println("\nTry and guess the code by providing input like: 'rrrr'")
	EnterAGuess()
}

func EnterAGuess() {
	fmt.Println("Enter your guess:")
	prompt()
}

func prompt() {
	fmt.Print("> ")
}
