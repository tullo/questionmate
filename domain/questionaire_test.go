package domain

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQuestionaire(t *testing.T) {
	fn := fmt.Sprintf("%s/src/github.com/rwirdemann/questionmate/config/surf.yaml", os.Getenv("GOPATH"))
	dat, err := ioutil.ReadFile(fn)
	assert.Nil(t, err)
	q := NewQuestionaire(dat)
	assert.Equal(t, "How is the weather today?", q.Questions[1].Text)
	assert.Equal(t, "Sunny", q.Questions[1].Options[1].Text)
	assert.Equal(t, 10, q.Questions[1].Options[1].Targets["cycling"])
	assert.Equal(t, "Rainy", q.Questions[1].Options[4].Text)
	assert.Equal(t, 8, q.Questions[1].Options[4].Targets["running"])
}
