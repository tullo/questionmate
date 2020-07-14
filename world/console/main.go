package main

import (
	"github.com/rwirdemann/questionmate/adapter/driver/console"
	"github.com/rwirdemann/questionmate/adapter/repositories/file"
	"github.com/rwirdemann/questionmate/domain"
	"github.com/rwirdemann/questionmate/usecase"
)

func main() {
	// 1. Instantiate the "I need to go out httpadapter"
	repositoryAdapter := file.NewQuestionRepository("legacylab.qm")

	// 2. Instantiate the hexagon
	hexagon := usecase.NextQuestion{QuestionRepository: repositoryAdapter}

	// 3. Instantiate the "I need to go in adapter"
	consoleAdapter := console.NewAdapter(hexagon)

	var answers []domain.Answer
	a, ok := consoleAdapter.Ask(answers)
	for ok {
		answers = append(answers, a)
		a, ok = consoleAdapter.Ask(answers)
	}
}
