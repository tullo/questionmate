package main

import (
	"fmt"
	"github.com/rwirdemann/questionmate/adapter/driver/console"
	"github.com/rwirdemann/questionmate/adapter/repositories/file"
	"github.com/rwirdemann/questionmate/domain"
	"github.com/rwirdemann/questionmate/usecase"
	"log"
)

func main() {
	// 1. Instantiate the "I need to go out httpadapter"
	repositoryAdapter := file.NewQuestionRepository("legacylab.qm")

	// 2. Instantiate the hexagon
	hexagon := usecase.NextQuestion{QuestionRepository: repositoryAdapter}

	// 3. Instantiate the "I need to go in adapter"
	consoleAdapter := console.NewAdapter(hexagon)

	var answers []domain.Answer
	q, hasNext := consoleAdapter(answers)
	for hasNext {
		fmt.Printf("%s\n", q.Text)
		for _, option := range q.Options {
			fmt.Printf("%d: %s\n", option.Value, option.Text)
		}
		fmt.Print("Your answer: ")
		var answer string
		_, err := fmt.Scanln(&answer)
		if err != nil {
			log.Fatal(err)
		}

		option, isValidAnswer := q.GetOptionByString(answer)
		for !isValidAnswer {
			fmt.Print("Try again: ")
			_, _ = fmt.Scanln(&answer)
			option, isValidAnswer = q.GetOptionByString(answer)
		}
		a := domain.Answer{QuestionID: q.ID, Value: option.Value}
		answers = append(answers, a)
		q, hasNext = consoleAdapter(answers)
	}
}
