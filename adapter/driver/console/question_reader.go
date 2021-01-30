package console

import (
	"fmt"
	"log"

	"github.com/tullo/questionmate/domain"
	"github.com/tullo/questionmate/usecase"
)

type Adapter struct {
	reader usecase.QuestionReader
}

func NewAdapter(questionReader usecase.QuestionReader) Adapter {
	return Adapter{reader: questionReader}
}

func (a Adapter) Ask(answers []domain.Answer) (domain.Answer, bool) {
	q, hasNext := a.reader.NextQuestion(answers)
	if hasNext {
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
		return domain.Answer{QuestionID: q.ID, Value: option.Value}, true
	}
	return domain.Answer{}, false
}
