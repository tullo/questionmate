package console

import (
	"github.com/rwirdemann/questionmate/domain"
	"github.com/rwirdemann/questionmate/usecase"
)

var reader usecase.QuestionReader

func NewAdapter(questionReader usecase.QuestionReader) func(answer []domain.Answer) (domain.Question, bool) {
	reader = questionReader
	return nextQuestion
}

func nextQuestion(answers []domain.Answer) (domain.Question, bool) {
	return reader.NextQuestion(answers)
}
