package ports

import "github.com/rwirdemann/questionmate/domain"

// QuestionRepository models the repository from where the questions are read
// from.
type QuestionRepository interface {
	LoadQuestions(data []byte) domain.Questionaire
}
