package qm

import (
	"github.com/rwirdemann/questionmate/domain"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type QuestionStore struct {
}

func (q QuestionStore) LoadQuestions(data []byte) domain.Questionaire {
	var questionaire domain.Questionaire
	questionaire.Questions = make(map[int]*domain.Question)
	lines := strings.Split(string(data), "\n")

	var question *domain.Question
	var option *domain.Option
	for _, l := range lines {
		if isQuestion(l) {
			q := strings.Split(l, ":")
			id, err := strconv.Atoi(q[0])
			if err != nil {
				log.Fatal(err)
			}
			question = domain.NewQuestion(id, strings.Trim(q[1], " "))
			questionaire.Questions[id] = question
		}

		if question != nil && isType(l) {
			t := strings.Split(l, ":")
			question.Type = strings.Trim(t[1], " ")
		}

		if question != nil && isOption(l) {
			o := strings.Split(l, ":")
			id, err := strconv.Atoi(o[0])
			if err != nil {
				log.Fatal(err)
			}
			option = &domain.Option{ID: id, Text: strings.Trim(o[1], " ")}
			question.Options[id] = option
		}
	}

	return questionaire
}

func isOption(s string) bool {
	match, _ := regexp.MatchString("(^ {2}\\d+):", s)
	return match
}

func isType(s string) bool {
	return strings.HasPrefix(s, "type:")
}

func isQuestion(s string) bool {
	match, _ := regexp.MatchString("(^\\d+):", s)
	return match
}
