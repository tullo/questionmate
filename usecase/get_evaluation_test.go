package usecase

import (
	"testing"

	"github.com/rwirdemann/questionmate/domain"
	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	q1 := domain.NewQuestion(1, "q1")
	o1 := domain.NewOption(1, "o1")
	o1.Targets["t1"] = domain.Score{Value: 10}
	q1.Options = append(q1.Options, o1)
	var questionRepository questionRepository = MockQuestionRepository{questions: []domain.Question{q1}}

	a1 := domain.Answer{
		QuestionID: 1,
		Value:      1,
	}
	answers := []domain.Answer{a1}
	usecase := Evaluator{QuestionRepository: questionRepository}
	e := usecase.GetEvaluation(answers)
	t1, ok := e.GetTarget("t1")
	assert.True(t, ok)
	assert.Equal(t, 10, t1.Score)
}
