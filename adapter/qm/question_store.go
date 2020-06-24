package qm

import (
	"github.com/rwirdemann/questionmate/domain"
)

type QuestionStore struct {
	filename string
}

func NewQuestionStore(filename string) QuestionStore {
	return QuestionStore{filename: filename}
}

func (q QuestionStore) LoadQuestions(data []byte) domain.Questionaire {
	var questionaire domain.Questionaire
	questionaire.Questions = make(map[int]domain.Question)
	return questionaire
}
