package nextquestion

import "github.com/rwirdemann/questionmate/domain"

type Answer struct {
}

type UseCase struct {
}

func (uc UseCase) NextQuestion(answers []Answer) domain.Question {
	return domain.Question{}
}
