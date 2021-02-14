package console

import (
	"fmt"
	"log"

	"github.com/tullo/questionmate/domain"
	"github.com/tullo/questionmate/usecase"
)

type Adapter struct {
	qr usecase.QuestionReader
}

func NewAdapter(r usecase.QuestionReader) Adapter {
	return Adapter{qr: r}
}

func (a Adapter) Ask(as []domain.Answer) (domain.Answer, bool) {
	q, hasNext := a.qr.NextQuestion(as)
	if hasNext {
		fmt.Printf("%s\n", q.Text)
		for _, o := range q.Options {
			fmt.Printf("%d: %s\n", o.Value, o.Text)
		}
		fmt.Print("Your answer: ")
		var answer string
		_, err := fmt.Scanln(&answer)
		if err != nil {
			log.Fatal(err)
		}

		o, ok := q.GetOptionByString(answer)
		for !ok {
			fmt.Print("Try again: ")
			_, _ = fmt.Scanln(&answer)
			o, ok = q.GetOptionByString(answer)
		}
		return domain.Answer{QuestionID: q.ID, Value: o.Value}, true
	}
	return domain.Answer{}, false
}
