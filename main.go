package main

import (
	"bufio"
	"fmt"
	"github.com/agomesd/go-pokedex/utils"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		userInput := scanner.Text()
		fmt.Printf("Pokedex > %s\n", userInput)
		cleaned := utils.CleanInput(userInput)
		fmt.Printf("Your command was: %s\n", cleaned[0])
	}
}
