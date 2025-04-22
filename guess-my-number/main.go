package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/andrey11/golang-experiment/guess-my-number/constants"
)

func main() {
	secretNumber := rand.Intn(100) + 1
	attempts := 0
	remainingAttempts := 10

	fmt.Println(constants.WelcomeMsg)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(constants.GuessMsg)
		scanner.Scan()

		input := strings.TrimSpace(scanner.Text())

		guess, err := tryGetNumber(input)
		
		if err != nil {
			fmt.Println(err)
			continue
		}

		attempts = incrementAttempts(attempts)
		remainingAttempts = decrementRemainingAttempts(remainingAttempts)

		if(remainingAttempts <= 0) {
			fmt.Printf(constants.GameOverMsg, secretNumber)
			break
		}
		
		if(!checkGuess(guess, secretNumber)) {
			fmt.Printf(constants.RemainingAttemptsMsg, remainingAttempts)
			continue
		}

		fmt.Printf(constants.SuccessMsg, attempts)

		break
	}

}

func tryGetNumber(input string) (int, error) {
	guess, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return 0, fmt.Errorf(constants.InvalidInputMsg)
	}
	return guess, nil
}

func checkGuess(guess int, secretNumber int) bool {
	if guess < secretNumber {
		fmt.Println(constants.TooLowMsg)
		return false
	}

	if guess > secretNumber {
		fmt.Println(constants.TooHighMsg)
		return false
	}

	fmt.Println(constants.CorrectMsg)
	return true
}

func incrementAttempts(attempts int) int {
	return attempts + 1
}

func decrementRemainingAttempts(remainingAttempts int) int {
	return remainingAttempts - 1
}
