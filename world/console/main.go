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
	dir, _ := os.Getwd()
	// 1. Instantiate the "I need to go out adapter"
	fn := fmt.Sprintf("%s/config/coma", dir)
	repo := file.NewQuestionnaireRepository(fn, parser.YAMLParser{})

	// 2. Instantiate the hexagon
	hexagon := usecase.NextQuestion{QR: repo}

	// 3. Instantiate the "I need to go in adapter"
	ca := console.NewAdapter(hexagon)

	var answers []domain.Answer
	a, ok := ca.Ask(answers)
	for ok {
		answers = append(answers, a)
		a, ok = ca.Ask(answers)
	}
}
