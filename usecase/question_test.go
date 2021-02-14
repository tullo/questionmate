package usecase

import (
	"github.com/stretchr/testify/assert"
	"github.com/tullo/questionmate/domain"

	"testing"
)

type MockQuestionRepository struct {
	qs      []domain.Question
	ratings map[string][]domain.Rating
}

func (MockQuestionRepository) GetQuestionnaire() domain.Questionnaire {
	panic("implement me")
}

func (mr MockQuestionRepository) GetRatings() map[string][]domain.Rating {
	return mr.ratings
}

func (mr MockQuestionRepository) GetDescriptions() map[int]string {
	return make(map[int]string)
}

func (mr MockQuestionRepository) GetQuestions() []domain.Question {
	return mr.qs
}

// The test is the adapter
func TestNextQuestion(t *testing.T) {
	var qs []domain.Question
	qs = append(qs, domain.Question{ID: 1})
	qs = append(qs, domain.Question{ID: 2})
	qs = append(qs, domain.Question{ID: 3})
	qs = append(qs, domain.Question{ID: 4})

	q5 := domain.NewQuestion(5, "5")
	q5.Dependencies[4] = 1
	qs = append(qs, q5)

	repo := MockQuestionRepository{qs: qs}
	hexagon := NextQuestion{QR: repo}
	var as []domain.Answer
	q, b := hexagon.NextQuestion(as)
	assert.True(t, b)
	assert.Equal(t, 1, q.ID)

	as = append(as, domain.Answer{QuestionID: 1})
	q, b = hexagon.NextQuestion(as)
	assert.True(t, b)
	assert.Equal(t, 2, q.ID)

	as = append(as, domain.Answer{QuestionID: 2})
	q, b = hexagon.NextQuestion(as)
	assert.True(t, b)
	assert.Equal(t, 3, q.ID)
}
