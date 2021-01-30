package main

import (
	"fmt"
	"os"

	"github.com/tullo/questionmate/adapter/driver/console"
	"github.com/tullo/questionmate/adapter/repositories/file"
	"github.com/tullo/questionmate/adapter/repositories/parser"
	"github.com/tullo/questionmate/domain"
	"github.com/tullo/questionmate/usecase"
)

func main() {
	// 1. Instantiate the "I need to go out adapter"
	fn := fmt.Sprintf("%s/src/github.com/tullo/questionmate/config/coma", os.Getenv("GOPATH"))
	repositoryAdapter := file.NewQuestionRepository(fn, parser.YAMLParser{})

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
