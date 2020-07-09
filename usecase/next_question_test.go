package usecase

import (
	"github.com/rwirdemann/questionmate/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Left side port
type QuestionReader interface {
	NextQuestion(answers []domain.Answer) (domain.Question, bool)
}

type MockQuestionRepository struct {
}

func (m MockQuestionRepository) GetQuestions() []domain.Question {
	var questions []domain.Question
	questions = append(questions, domain.Question{ID: 1})
	questions = append(questions, domain.Question{ID: 2})
	return questions
}

// The test is the adapter
func TestNextQuestion(t *testing.T) {
	var questionRepository QuestionRepository = MockQuestionRepository{}
	var questionReader QuestionReader = NextQuestion{QuestionRepository: questionRepository}
	var answers []domain.Answer
	q, b := questionReader.NextQuestion(answers)
	assert.True(t, b)
	assert.Equal(t, 1, q.ID)

	answers = append(answers, domain.Answer{QuestionID: 1})
	q, b = questionReader.NextQuestion(answers)
	assert.True(t, b)
	assert.Equal(t, 2, q.ID)

	answers = append(answers, domain.Answer{QuestionID: 2})
	_, b = questionReader.NextQuestion(answers)
	assert.False(t, b)
}
