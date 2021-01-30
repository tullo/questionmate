package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tullo/questionmate/domain"
)

func TestProcess(t *testing.T) {
	q1 := domain.NewQuestion(1, "q1")
	o1 := domain.NewOption(1, "o1")
	o1.Targets["t1"] = domain.Score{Value: 10}
	q1.Options = append(q1.Options, o1)

	r1 := domain.Rating{
		Target:      "t1",
		Min:         0,
		Max:         10,
		Description: "This is really good",
	}
	ratings := make(map[string][]domain.Rating)
	ratings["t1"] = []domain.Rating{r1}

	var questionRepository questionRepository = MockQuestionRepository{
		questions: []domain.Question{q1},
		ratings:   ratings,
	}

	a1 := domain.Answer{
		QuestionID: 1,
		Value:      1,
	}
	answers := []domain.Answer{a1}
	usecase := Assessment{QuestionRepository: questionRepository}
	e := usecase.GetAssessment(answers)
	t1, ok := e.GetTarget("t1")
	assert.True(t, ok)
	assert.Equal(t, 10, t1.Score)
	assert.Equal(t, "This is really good", t1.Rating)
}
