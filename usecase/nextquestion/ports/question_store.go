package ports

import "github.com/rwirdemann/questionmate/domain"

type QuestionStore interface {
	LoadQuestions(data []byte) domain.Questionaire
}
