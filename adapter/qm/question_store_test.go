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
	fn := fmt.Sprintf("%s/src/github.com/rwirdemann/questionmate/config/%s", os.Getenv("GOPATH"), "legacylab.qm")
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	var store QuestionStore
	q := store.LoadQuestions(data)
	assert.Equal(t, "Schätzen Sie relativen Marktwert der Software auf einer Skala zwischen 0 und 100", q.Questions[10].Text)
}
