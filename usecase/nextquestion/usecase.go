package nextquestion

import "github.com/rwirdemann/questionmate/domain"

type UseCase struct {
}

func (uc UseCase) NextQuestion(answers []domain.Answer) domain.Question {
	return domain.Question{}
}
