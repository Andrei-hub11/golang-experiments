package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/andrey11/golang-experiment/quiz/constants"
	"github.com/andrey11/golang-experiment/quiz/internal"
	"github.com/andrey11/golang-experiment/quiz/models"
	"github.com/fatih/color"
)

func main() {
	currentQuestion := 0
	totalQuestions := 0
	totalCorrectAnswers := 0
	totalIncorrectAnswers := 0

	questions, err := internal.LoadQuestions("questions.json")
	if err != nil {	
		fmt.Println("Error loading questions:", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(constants.WelcomeMessage)	
	totalQuestions = getTotalQuestions(questions)
	fmt.Printf(constants.TotalQuestionsMessage, totalQuestions)

	for currentQuestion < totalQuestions {
		internal.ShuffleOptions(&questions[currentQuestion])
		printQuestion(questions[currentQuestion])

		mapQuestion := createMapToCurrentQuestion(questions[currentQuestion])
		currentAnswer := getAnswer(mapQuestion, questions[currentQuestion])

		fmt.Println(constants.GameQuestion)
		userAnswer := getUserAnswer(scanner)

		if verifyAnswer(mapQuestion, userAnswer, questions[currentQuestion]) {
			color.Green(constants.CorrectAnswer)
			totalCorrectAnswers++
		} else {
			color.Red(constants.IncorrectAnswer, currentAnswer)
			totalIncorrectAnswers++
		}

		currentQuestion = incrementIndexQuestion(currentQuestion)
	}

	summarizeResults(totalCorrectAnswers, totalIncorrectAnswers)
}

func printQuestion(question models.Question) {
	fmt.Println(question.Question)
	for i, option := range question.Options {
		fmt.Printf("%d. %s\n", i + 1, option)
	}
}

func getAnswer(mapQuestion map[int]int, question models.Question) int {
	for key, value := range mapQuestion {
		if value == question.Answer {
			return key
		}
	}
	
	return 0
}

func getUserAnswer(scanner *bufio.Scanner) int {
for {
		scanner.Scan()
		input := scanner.Text()
		userAnswer, err := tryGetNumber(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

	if userAnswer < 1 || userAnswer > 4 {
		fmt.Println(constants.InputOutOfRangeMsg)
			continue
		}
		
		return userAnswer
	}
}

func getTotalQuestions(questions []models.Question) int {
	return len(questions)
}

func incrementIndexQuestion(index int) int {
	return index + 1
}

func createMapToCurrentQuestion(question models.Question) map[int]int {
	questionMap := make(map[int]int)
	for i := range question.Options {
		questionMap[i + 1] = i
	}
	
	return questionMap
}

func verifyAnswer(mapQuestion map[int]int, userAnswer int, currentQuestion models.Question) bool {
	return mapQuestion[userAnswer] == currentQuestion.Answer
}

func tryGetNumber(input string) (int, error) {
	guess, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return 0, fmt.Errorf(constants.InvalidInputMsg)
	}
	return guess, nil
}

func summarizeResults(totalCorrectAnswers int, totalIncorrectAnswers int) {
	fmt.Printf(constants.TotalCorrectAnswersMessage, totalCorrectAnswers)
	fmt.Printf(constants.TotalIncorrectAnswersMessage, totalIncorrectAnswers)
}
