package print

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func WelcomeMessage() {
	clearScreen()
	fmt.Println("Welcome to Mastermind!")
}

func QuitMessage() {
	for i := 0; i < 30; i++ {
		fmt.Println("byeeee")
		time.Sleep(30 * time.Millisecond)
	}
	clearScreen()
}

func Instructions() {
	clearScreen()
	fmt.Println("The game will generate a sequence with four elements made up of: ")
	fmt.Println("(r)ed, (g)reen, (b)lue, and (y)ellow. It will look like 'rybg'.")
	fmt.Println("Your goal is to guess the code in the fewest guesses possible.")
	fmt.Println("\nGuesses must be four characters long and only contain 'r', 'g', 'b' and 'y'.")
	fmt.Println("It doesn't matter if the guess is upper case or lower case.")
	fmt.Println("The game will then tell you correct positions and correct colors.")
	fmt.Println("Good luck!\n")
}

func UnknownCommand() {
	fmt.Println("Unknown command! Please try again ...\n")
}

func WhatsNext() {
	fmt.Println("Would you like to (p)lay read the (i)instructions or (q)uit?")
	prompt()
}

func GameStart() {
	clearScreen()
	fmt.Println("I have generated a random code four characters in length using the letters:")
	fmt.Println("'r', 'g', 'b' and 'y'")
	fmt.Println("\nTry and guess the code by providing input like: 'rrrr'")
}

func EnterAGuess() {
	fmt.Println("Enter your guess:")
	prompt()
}

func ThanksForPlaying() {
	fmt.Println("Thanks for playing!")
	fmt.Println("Remember, winners don't do drugs.\n")
}

func prompt() {
	fmt.Print("> ")
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
