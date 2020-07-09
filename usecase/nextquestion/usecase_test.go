package nextquestion

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/rwirdemann/questionmate/adapter/repositories/file"
	"github.com/rwirdemann/questionmate/domain"
	"github.com/stretchr/testify/assert"
)

var usecase UseCase

func init() {
	fn := fmt.Sprintf("%s/src/github.com/rwirdemann/questionmate/config/%s", os.Getenv("GOPATH"), "questionmate.qm")
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	store := file.QuestionRepository{}
	usecase = NewUseCase(store, data)
}

func TestNextQuestion(t *testing.T) {
	var answers []domain.Answer
	q, found := usecase.NextQuestion(answers)
	assert.True(t, found)
	assert.Equal(t, 10, q.ID)

	answers = append(answers, domain.Answer{QuestionID: 10})
	q, found = usecase.NextQuestion(answers)
	assert.True(t, found)
	assert.Equal(t, 20, q.ID)

	answers = append(answers, domain.Answer{QuestionID: 20})
	q, found = usecase.NextQuestion(answers)
	assert.True(t, found)
	assert.Equal(t, 40, q.ID)

	answers = append(answers, domain.Answer{QuestionID: 40})
	q, found = usecase.NextQuestion(answers)
	assert.True(t, found)
	assert.Equal(t, 44, q.ID)

	answers = append(answers, domain.Answer{QuestionID: 44})
	_, found = usecase.NextQuestion(answers)
	assert.False(t, found)
}
