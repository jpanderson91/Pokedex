package main

import (
	"fmt"
	"bufio"
	"os"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	//infinite for loop. This lool will execute once for every command the user types in (we don't want to exit the program after just one command)
	for {
		fmt.Print("Pokedex >")
		// Use the scanner's .Scan and .Text methods to get the user's input as a string
		scanner.Scan()
		input := scanner.Text()
		// clean the user's input string
		cleanedInput := cleanInput(input)
		// capture the first "word" of the input and use it to print: Your command was : <first word>
		if len(cleanedInput) > 0 {
			firstWord := cleanedInput[0]
			fmt.Printf("Your command was: %s\n", firstWord)
		}
	}


}



