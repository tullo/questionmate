package yaml

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadQuestions(t *testing.T) {
	fn := fmt.Sprintf("%s/src/github.com/rwirdemann/questionmate/config/%s", os.Getenv("GOPATH"), "boardbuddy.yaml")
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	var store QuestionStore
	q := store.LoadQuestions(data)
	assert.Equal(t, "freeride", q.Targets[1].Label)
	assert.Equal(t, "Welches sind die Hauptbedingungen, unter denen du surfst?", q.Questions[1].Text)
	fmt.Printf("%v", q.Questions[1].Options[1].Scores)
	assert.Equal(t, 10, q.Questions[1].Options[1].Scores[1].Value)
	assert.Equal(t, "Auf dem Meer kannst die lange Schläge fahren", q.Questions[1].Options[1].Scores[1].Why)
}
