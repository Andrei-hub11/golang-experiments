package internal

import (
	"math/rand"

	"github.com/andrey11/golang-experiment/quiz/models"
)

func ShuffleOptions(q *models.Question) {

	correctAnswer := q.Answer

	type optionWithIndex struct {
		Text  string
		Index int
	}

	options := make([]optionWithIndex, len(q.Options))
	for i, opt := range q.Options {
		options[i] = optionWithIndex{Text: opt, Index: i}
	}

	rand.Shuffle(len(options), func(i, j int) {
		options[i], options[j] = options[j], options[i]
	})

	newOptions := make([]string, len(q.Options))
	for i, opt := range options {
		newOptions[i] = opt.Text
		if opt.Index == correctAnswer {
			q.Answer = i 
		}
	}

	q.Options = newOptions
}
