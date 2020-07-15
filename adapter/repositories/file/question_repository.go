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
	Questions    []domain.Question
	Descriptions map[int]string
}

func NewQuestionRepository(file string) QuestionRepository {
	questions := repositories.ParseQuestions(readFile(file + ".questions"))
	descriptions := repositories.ParseDescriptions(readFile(file + ".desc"))
	return QuestionRepository{Questions: questions, Descriptions: descriptions}
}

func readFile(file string) []byte {
	fn := fmt.Sprintf("%s/src/github.com/rwirdemann/questionmate/config/%s", os.Getenv("GOPATH"), file)
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (q QuestionRepository) GetQuestions() []domain.Question {
	return q.Questions
}

func (q QuestionRepository) GetDescriptions() map[int]string {
	return q.Descriptions
}

func (q QuestionRepository) String() string {
	var s string
	for _, question := range q.Questions {
		s = fmt.Sprintf("%s---------------------------------------------------------------------------------------\n", s)
		s = fmt.Sprintf("%sFrage: %s\n", s, question.Text)
		if desc, ok := q.Descriptions[question.ID]; ok {
			s = fmt.Sprintf("%sBeschreibung: %s\n", s, desc)
		}
		for _, option := range question.Options {
			s = fmt.Sprintf("%s- %s\n", s, option.Text)
		}
	}
	return s
}
