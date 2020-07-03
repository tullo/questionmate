package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnanswered(t *testing.T) {
	q := NewQuestionnaire()
	q.Questions[1] = &Question{ID: 1}
	q.Questions[2] = &Question{ID: 2}

	var answers []Answer
	assert.Equal(t, []int{1, 2}, q.Unanswered(answers))

	answers = append(answers, Answer{ID: 1})
	assert.Equal(t, []int{2}, q.Unanswered(answers))
	answers = append(answers, Answer{ID: 2})

	var expected []int
	assert.Equal(t, expected, q.Unanswered(answers))
}
