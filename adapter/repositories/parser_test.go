package repositories

import (
	"fmt"
	"github.com/rwirdemann/questionmate/domain"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDescriptions(t *testing.T) {
	fn := fmt.Sprintf("%s/src/github.com/rwirdemann/questionmate/config/%s", os.Getenv("GOPATH"), "questionmate.desc")
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	descriptions := ParseDescriptions(data)
	assert.Equal(t, "This is a detailed description of the question that can be requested by clients on demand.", descriptions[10])
}

func TestParseTargets(t *testing.T) {
	fn := fmt.Sprintf("%s/src/github.com/rwirdemann/questionmate/config/%s", os.Getenv("GOPATH"), "questionmate.targets")
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	targets := ParseTargets(data)
	assert.Equal(t, "The target businessvalue evaluates a software in terms of its market value for the underlying business.", targets["businessvalue"])
}

func TestParseQuestions(t *testing.T) {
	fn := fmt.Sprintf("%s/src/github.com/rwirdemann/questionmate/config/%s", os.Getenv("GOPATH"), "questionmate.qm")
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	questions := ParseQuestions(data)
	q := questions[0]
	assert.Equal(t, "Estimate the proportional market value of your software on a range between 0 and 100.", q.Text)
	assert.Equal(t, "single", q.Type)
	assert.Equal(t, "This is a detailed description of the question that can be requested by clients on demand.", q.Desc)
	assert.Len(t, q.Options, 3)
	assert.Equal(t, "0 - 30", GetOption(q, 1).Text)

	// Targets
	assert.Len(t, GetOption(q, 1).Targets, 1)
	assert.Equal(t, 1, GetOption(q, 1).Targets["businessvalue"].Value)
	assert.Equal(t, 2, GetOption(q, 2).Targets["businessvalue"].Value)
	assert.Equal(t, 3, GetOption(q, 4).Targets["businessvalue"].Value)

	q = questions[1]
	assert.Equal(t, "How many features does your team develop per month?", q.Text)
	assert.Len(t, q.Options, 3)
	assert.Equal(t, "1", GetOption(q, 1).Text)

	// Targets
	assert.Len(t, q.Options[1].Targets, 2)
	assert.Equal(t, 1, GetOption(q, 1).Targets["extendability"].Value)
	assert.Equal(t, 1, GetOption(q, 1).Targets["businessvalue"].Value)

	// Dependencies
	q = questions[3]
	assert.Equal(t, 1, q.Dependencies[40])
}

func GetOption(q domain.Question, i int) domain.Option {
	o, _ := q.GetOption(i)
	return o
}

func Test_isQuestion(t *testing.T) {
	assert.True(t, isQuestion("10: Estimate the proportional"))
	assert.False(t, isQuestion("Estimate the proportional"))
}

func Test_isOption(t *testing.T) {
	assert.True(t, isOption("  1: 0 - 30"))
	assert.False(t, isOption("1: 0 - 30"))
}

func Test_isDependency(t *testing.T) {
	assert.True(t, isDependency("  40 => 1"))
	assert.False(t, isDependency("  40 = 1"))
	assert.False(t, isDependency(" 40 => 1"))
	assert.False(t, isDependency("  40 => a"))
}

func Test_isTarget(t *testing.T) {
	assert.True(t, isTarget("  - businessvalue: 1"))
	assert.False(t, isTarget(" - businessvalue: 1"))
	assert.False(t, isTarget("  - businessvalue: x"))
	assert.True(t, isTarget("  - businessvalue: 12"))
	assert.False(t, isTarget("  - businessvalue 12"))
}
