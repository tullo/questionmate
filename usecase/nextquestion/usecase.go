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
// has given so far. It returns false when there is no next question left.
func (uc UseCase) NextQuestion(answers []domain.Answer) (domain.Question, bool) {
	questionnaire := uc.questionStore.LoadQuestions(uc.data)
	unanswered := questionnaire.Unanswered(answers)
	if len(unanswered) > 0 {
		return *questionnaire.Questions[unanswered[0]], true
	}
	return domain.Question{}, false
}
