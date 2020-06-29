package nextquestion

import (
	"fmt"
	"github.com/rwirdemann/questionmate/adapter/qm"
	"github.com/rwirdemann/questionmate/domain"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var usecase UseCase

func init() {
	fn := fmt.Sprintf("%s/src/github.com/rwirdemann/questionmate/config/%s", os.Getenv("GOPATH"), "questionmate.qm")
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	store := qm.QuestionStore{}
	usecase = NewUseCase(store, data)
}

func TestNextQuestion(t *testing.T) {
	var answers []domain.Answer
	q := usecase.NextQuestion(answers)
	assert.Equal(t, 10, q.ID)

	answers = append(answers, domain.Answer{ID: 10})
	q = usecase.NextQuestion(answers)
	assert.Equal(t, 20, q.ID)
}
