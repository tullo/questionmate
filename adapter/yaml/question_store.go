package yaml

import (
	"github.com/rwirdemann/questionmate/domain"
	"gopkg.in/yaml.v2"
	"log"
)

type QuestionStore struct {
	filename string
}

func NewQuestionStore(filename string) QuestionStore {
	return QuestionStore{filename: filename}
}

func (q QuestionStore) LoadQuestions(data []byte) domain.Questionaire {
	var questionaire domain.Questionaire
	questionaire.Questions = make(map[int]*domain.Question)
	err := yaml.Unmarshal(data, &questionaire)
	if err != nil {
		log.Fatal(err)
	}
	return questionaire
}
