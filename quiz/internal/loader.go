package internal

import (
	"encoding/json"
	"os"

	"github.com/andrey11/golang-experiment/quiz/models"
)

func LoadQuestions(filepath string) ([]models.Question, error) {
	jsonFile, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var questions []models.Question
	
	err = json.Unmarshal(jsonFile, &questions)

	if err != nil {
		return nil, err
	}

	return questions, nil
}

