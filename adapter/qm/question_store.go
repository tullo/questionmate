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
	questionaire.Questions = make(map[int]*domain.Question)
	lines := strings.Split(string(data), "\n")

	var question *domain.Question
	for _, l := range lines {
		if question == nil && isQuestion(l) {
			q := strings.Split(l, ":")
			id, err := strconv.Atoi(q[0])
			if err != nil {
				log.Fatal(err)
			}
			question = &domain.Question{ID: id, Text: strings.Trim(q[1], " ")}
			questionaire.Questions[id] = question
		}
		if question != nil && isType(l) {
			t := strings.Split(l, ":")
			question.Type = strings.Trim(t[1], " ")
		}
	}

	return questionaire
}

func isType(s string) bool {
	return strings.HasPrefix(s, "type:")
}

func isQuestion(s string) bool {
	match, _ := regexp.MatchString("(^\\d+):", s)
	return match
}
