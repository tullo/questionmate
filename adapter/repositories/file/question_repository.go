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
	Targets      map[string]string
}

func NewQuestionRepository(file string) QuestionRepository {
	var questions []domain.Question
	var descriptions map[int]string
	var targets map[string]string

	if bytes, ok := readFile(file + ".questions"); ok {
		questions = repositories.ParseQuestions(bytes)
	}
	if bytes, ok := readFile(file + ".desc"); ok {
		descriptions = repositories.ParseDescriptions(bytes)
	}
	if bytes, ok := readFile(file + ".targets"); ok {
		targets = repositories.ParseTargets(bytes)
	}
	return QuestionRepository{Questions: questions, Descriptions: descriptions, Targets: targets}
}

func readFile(file string) ([]byte, bool) {
	fn := fmt.Sprintf("%s/src/github.com/rwirdemann/questionmate/config/%s", os.Getenv("GOPATH"), file)
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Println(err)
		return nil, false
	}
	return data, true
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
