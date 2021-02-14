package usecase

import (
	"sort"

	"github.com/tullo/questionmate/domain"
)

// Left side port
type QuestionReader interface {
	NextQuestion(as []domain.Answer) (domain.Question, bool)
}

// Right side port
type questionRepository interface {
	GetQuestionnaire() domain.Questionnaire
	GetQuestions() []domain.Question
	GetRatings() map[string][]domain.Rating
	GetDescriptions() map[int]string
}

// Hexagon
type NextQuestion struct {
	QR questionRepository
}

func (nc NextQuestion) NextQuestion(as []domain.Answer) (domain.Question, bool) {
	qs := nc.QR.GetQuestions()
	unanswered := unanswered(as, qs)
	if len(unanswered) > 0 {
		q := byID(unanswered[0], qs)
		if desc, ok := nc.QR.GetDescriptions()[q.ID]; ok {
			q.Desc = desc
		}
		return q, true
	}
	return domain.Question{}, false
}

func byID(id int, qs []domain.Question) domain.Question {
	for _, q := range qs {
		if q.ID == id {
			return q
		}
	}
	return domain.Question{}
}

// unanswered returns an sorted array of unanswered question ids according to
// the given answers.
func unanswered(as []domain.Answer, qs []domain.Question) []int {
	var unanswered []int
	for _, q := range qs {
		if !contains(q.ID, as) {
			unanswered = append(unanswered, q.ID)
		}
	}
	sort.Ints(unanswered)
	return unanswered
}

func contains(id int, as []domain.Answer) bool {
	for _, a := range as {
		if a.QuestionID == id {
			return true
		}
	}
	return false
}
