package parser

import "github.com/tullo/questionmate/domain"

type Parser interface {
	ParseQuestionnaire(data []byte) domain.Questionnaire
	ParseQuestions(data []byte) []domain.Question
	ParseRatings(data []byte) map[string][]domain.Rating
	Suffix() string
}
