package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseYAMLQuestions(t *testing.T) {
	fn := fmt.Sprintf("%s/src/github.com/rwirdemann/questionmate/config/%s", os.Getenv("GOPATH"), "coma/questions.yaml")
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	parser := YAMLParser{}
	questions := parser.ParseQuestions(data)
	assert.Equal(t, 1, questions[0].ID)
	assert.Equal(t, "Die aktuellen Sourcen sind...", questions[0].Text)
}
