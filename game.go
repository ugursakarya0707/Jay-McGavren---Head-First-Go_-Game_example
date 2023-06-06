package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	targetInt := rand.Intn(100)
	fmt.Println(targetInt)
	fmt.Println("Guess what the target number is?")

	for guessses := 0; guessses < 10; guessses++ {
		fmt.Println("You have left", 10-guessses, "guesses")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("your guess is :", input)
		guess, err := strconv.Atoi(input)

		if guess < targetInt {
			fmt.Println("Oops your guess was low")
		} else if guess > targetInt {
			fmt.Println("Oops your guess was high")
		} else {
			fmt.Println("You made it")
			break
		}

	}

}
