package file

import (
	"fmt"
	"github.com/rwirdemann/questionmate/adapter/repositories/parser"
	"github.com/rwirdemann/questionmate/domain"
	"io/ioutil"
	"log"
)

type QuestionRepository struct {
	Questions    []domain.Question
	Descriptions map[int]string
	Targets      map[string]string
}

func NewQuestionRepository(path string, parser parser.Parser) QuestionRepository {
	var questions []domain.Question
	var descriptions map[int]string
	var targets map[string]string

	if bytes, ok := readFile(path + "/questions." + parser.Suffix()); ok {
		questions = parser.ParseQuestions(bytes)
	}
	if bytes, ok := readFile(path + "/descriptions." + parser.Suffix()); ok {
		descriptions = parser.ParseDescriptions(bytes)
	}
	if bytes, ok := readFile(path + "/targets" + parser.Suffix()); ok {
		targets = parser.ParseTargets(bytes)
	}
	return QuestionRepository{Questions: questions, Descriptions: descriptions, Targets: targets}
}

func readFile(file string) ([]byte, bool) {
	data, err := ioutil.ReadFile(file)
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
			for k, v := range option.Targets {
				s = fmt.Sprintf("%s  - %s: %d\n", s, k, v.Value)
			}
		}
	}
	return s
}
