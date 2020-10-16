package usecase

import "github.com/rwirdemann/questionmate/domain"

type Questionnaire struct {
	Repository questionRepository
}

func (receiver Questionnaire) Get() domain.Questionnaire {
	return receiver.Repository.GetQuestionnaire()
}
