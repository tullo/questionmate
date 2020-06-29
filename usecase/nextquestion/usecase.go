package nextquestion

import (
	"github.com/rwirdemann/questionmate/domain"
	"github.com/rwirdemann/questionmate/usecase/nextquestion/ports"
)

type UseCase struct {
	questionStore ports.QuestionStore
	data          []byte
}

func NewUseCase(store ports.QuestionStore, data []byte) UseCase {
	return UseCase{questionStore: store, data: data}
}

// NextQuestion returns the next questions according to the answers the user
// has given so far.
func (uc UseCase) NextQuestion(answers []domain.Answer) domain.Question {
	questionnaire := uc.questionStore.LoadQuestions(uc.data)
	unanswered := questionnaire.Unanswered(answers)
	return *questionnaire.Questions[unanswered[0]]
}
