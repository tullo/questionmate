package parser

import "github.com/rwirdemann/questionmate/domain"

type Parser interface {
	ParseQuestions(data []byte) []domain.Question
	ParseDescriptions(data []byte) map[int]string
	ParseTargets(data []byte) map[string]string
	Suffix() string
}
