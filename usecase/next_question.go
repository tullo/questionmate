package usecase

import (
	"github.com/rwirdemann/questionmate/domain"
	"sort"
)

// Left side port
type QuestionReader interface {
	NextQuestion(answers []domain.Answer) (domain.Question, bool)
}

// Right side port
type questionRepository interface {
	GetQuestions() []domain.Question
	GetRatings() map[string][]domain.Rating
	GetDescriptions() map[int]string
}

// Hexagon
type NextQuestion struct {
	QuestionRepository questionRepository
}

func (nc NextQuestion) NextQuestion(answers []domain.Answer) (domain.Question, bool) {
	questions := nc.QuestionRepository.GetQuestions()
	unanswered := unanswered(answers, questions)
	if len(unanswered) > 0 {
		question := byID(unanswered[0], questions)
		if desc, ok := nc.QuestionRepository.GetDescriptions()[question.ID]; ok {
			question.Desc = desc
		}
		return question, true
	}
	return domain.Question{}, false
}

func byID(id int, questions []domain.Question) domain.Question {
	for _, q := range questions {
		if q.ID == id {
			return q
		}
	}
	return domain.Question{}
}

// unanswered returns an sorted array of unanswered question ids according to
// the given answers.
func unanswered(answers []domain.Answer, questions []domain.Question) []int {
	var unanswered []int
	for _, q := range questions {
		if !contains(q.ID, answers) {
			unanswered = append(unanswered, q.ID)
		}
	}
	sort.Ints(unanswered)
	return unanswered
}

func contains(id int, answers []domain.Answer) bool {
	for _, a := range answers {
		if a.QuestionID == id {
			return true
		}
	}
	return false
}
