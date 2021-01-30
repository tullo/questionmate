package usecase

import (
	"github.com/stretchr/testify/assert"
	"github.com/tullo/questionmate/domain"

	"testing"
)

type MockQuestionRepository struct {
	questions []domain.Question
	ratings   map[string][]domain.Rating
}

func (m MockQuestionRepository) GetQuestionnaire() domain.Questionnaire {
	panic("implement me")
}

func (m MockQuestionRepository) GetRatings() map[string][]domain.Rating {
	return m.ratings
}

func (m MockQuestionRepository) GetDescriptions() map[int]string {
	return make(map[int]string)
}

func (m MockQuestionRepository) GetQuestions() []domain.Question {
	return m.questions
}

// The test is the adapter
func TestNextQuestion(t *testing.T) {
	var questions []domain.Question
	questions = append(questions, domain.Question{ID: 1})
	questions = append(questions, domain.Question{ID: 2})
	questions = append(questions, domain.Question{ID: 3})
	questions = append(questions, domain.Question{ID: 4})

	q5 := domain.NewQuestion(5, "5")
	q5.Dependencies[4] = 1
	questions = append(questions, q5)

	var questionRepository questionRepository = MockQuestionRepository{questions: questions}
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
}
