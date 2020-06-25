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
}

func Test_isQuestion(t *testing.T) {
	assert.True(t, isQuestion("10: Estimate the proportional"))
	assert.False(t, isQuestion("Estimate the proportional"))
}
