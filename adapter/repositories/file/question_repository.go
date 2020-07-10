package file

import (
	"fmt"
	"github.com/rwirdemann/questionmate/adapter/repositories"
	"io/ioutil"
	"os"

	"github.com/rwirdemann/questionmate/domain"
	"log"
)

type QuestionRepository struct {
	Questions []domain.Question
}

func NewQuestionRepository(file string) QuestionRepository {
	fn := fmt.Sprintf("%s/src/github.com/rwirdemann/questionmate/config/%s", os.Getenv("GOPATH"), file)
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	return QuestionRepository{Questions: repositories.ParseQuestions(data)}
}

func (q QuestionRepository) GetQuestions() []domain.Question {
	return q.Questions
}
