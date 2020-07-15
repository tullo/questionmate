package usecase

import (
	"github.com/rwirdemann/questionmate/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockQuestionRepository struct {
}

func (m MockQuestionRepository) GetDescriptions() map[int]string {
	return make(map[int]string)
}

func (m MockQuestionRepository) GetQuestions() []domain.Question {
	var questions []domain.Question
	questions = append(questions, domain.Question{ID: 1})
	questions = append(questions, domain.Question{ID: 2})
	questions = append(questions, domain.Question{ID: 3})
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
	q, b = questionReader.NextQuestion(answers)
	assert.True(t, b)
	assert.Equal(t, 3, q.ID)

	answers = append(answers, domain.Answer{QuestionID: 3})
	_, b = questionReader.NextQuestion(answers)
	assert.False(t, b)
}
