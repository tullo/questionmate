package qm

import (
	"github.com/rwirdemann/questionmate/domain"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type QuestionStore struct {
	filename string
}

func NewQuestionStore(filename string) QuestionStore {
	return QuestionStore{filename: filename}
}

func (q QuestionStore) LoadQuestions(data []byte) domain.Questionaire {
	var questionaire domain.Questionaire
	questionaire.Questions = make(map[int]domain.Question)
	lines := strings.Split(string(data), "\n")
	for _, l := range lines {
		if isQuestion(l) {
			q := strings.Split(l, ":")
			id, err := strconv.Atoi(q[0])
			if err != nil {
				log.Fatal(err)
			}
			questionaire.Questions[id] = domain.Question{ID: id, Text: strings.Trim(q[1], " ")}
		}
	}

	return questionaire
}

func isQuestion(s string) bool {
	match, _ := regexp.MatchString("(^\\d+):", s)
	return match
}
