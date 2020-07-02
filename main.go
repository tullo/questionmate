package main

import (
	"fmt"
	"github.com/rwirdemann/questionmate/adapter/qm"
	"github.com/rwirdemann/questionmate/domain"
	"github.com/rwirdemann/questionmate/usecase/nextquestion"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	fn := fmt.Sprintf("%s/src/github.com/rwirdemann/questionmate/config/%s", os.Getenv("GOPATH"), "legacylab.qm")
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	var store qm.QuestionStore
	uc := nextquestion.NewUseCase(store, data)
	var answers []domain.Answer
	hasNext := true
	for hasNext {
		var nextQuestion domain.Question
		nextQuestion, hasNext = uc.NextQuestion(answers)
		if hasNext {
			fmt.Printf("%s\n", nextQuestion.Text)
			for _, option := range nextQuestion.Options {
				fmt.Printf("%d: %s\n", option.ID, option.Text)
			}
			fmt.Print("Your answer: ")
			var answer string
			_, err := fmt.Scanln(&answer)
			if err != nil {
				log.Fatal(err)
			}

			isValidOption := false
			for !isValidOption {
				_, isValidOption = validateOption(answer, nextQuestion.Options)
				if !isValidOption {
					fmt.Print("Try again: ")
					_, err := fmt.Scanln(&answer)
					if err != nil {
						log.Fatal(err)
					}
				} else {
					a := domain.Answer{ID: nextQuestion.ID}
					answers = append(answers, a)
				}
			}
		}
	}
}

func validateOption(s string, options map[int]*domain.Option) (int, bool) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1, false
	}
	if _, ok := options[i]; ok {
		return i, true
	}
	return -1, false
}
