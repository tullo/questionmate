package parser

import "github.com/rwirdemann/questionmate/domain"

type Parser interface {
	ParseQuestions(data []byte) []domain.Question
	Suffix() string
}
