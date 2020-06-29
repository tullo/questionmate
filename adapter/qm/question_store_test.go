package qm

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadQuestions(t *testing.T) {
	fn := fmt.Sprintf("%s/src/github.com/rwirdemann/questionmate/config/%s", os.Getenv("GOPATH"), "questionmate.qm")
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	var store QuestionStore
	q := store.LoadQuestions(data)
	assert.Equal(t, "Estimate the proportional market value of you software on a range between 0 and 100.", q.Questions[10].Text)
	assert.Equal(t, "single", q.Questions[10].Type)
	assert.Len(t, q.Questions[10].Options, 3)
	assert.Equal(t, "0 - 30", q.Questions[10].Options[1].Text)

	// Targets
	assert.Len(t, q.Questions[10].Options[1].Targets, 1)
	assert.Equal(t, 1, q.Questions[10].Options[1].Targets["businessvalue"].Value)
	assert.Equal(t, 2, q.Questions[10].Options[2].Targets["businessvalue"].Value)
	assert.Equal(t, 3, q.Questions[10].Options[4].Targets["businessvalue"].Value)

	assert.Equal(t, "How many features does your team develop per month?", q.Questions[20].Text)
	assert.Len(t, q.Questions[20].Options, 3)
	assert.Equal(t, "1", q.Questions[20].Options[1].Text)

	// Targets
	assert.Len(t, q.Questions[20].Options[1].Targets, 2)
	assert.Equal(t, 1, q.Questions[20].Options[1].Targets["extendability"].Value)
	assert.Equal(t, 1, q.Questions[20].Options[1].Targets["businessvalue"].Value)
}

func Test_isQuestion(t *testing.T) {
	assert.True(t, isQuestion("10: Estimate the proportional"))
	assert.False(t, isQuestion("Estimate the proportional"))
}

func Test_isOption(t *testing.T) {
	assert.True(t, isOption("  1: 0 - 30"))
	assert.False(t, isOption("1: 0 - 30"))
}

func Test_isTarget(t *testing.T) {
	assert.True(t, isTarget("  - businessvalue: 1"))
	assert.False(t, isTarget(" - businessvalue: 1"))
	assert.False(t, isTarget("  - businessvalue: x"))
	assert.True(t, isTarget("  - businessvalue: 12"))
	assert.False(t, isTarget("  - businessvalue 12"))
}
