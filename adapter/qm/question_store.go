package qm

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/rwirdemann/questionmate/domain"
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
			id := toInt(q[0])
			question = domain.NewQuestion(id, strings.Trim(q[1], " "))
			questionaire.Questions[id] = question
		}

		if question != nil && isType(l) {
			t := strings.Split(l, ":")
			question.Type = strings.Trim(t[1], " ")
		}

		if question != nil && isOption(l) {
			o := strings.Split(l, ":")
			id := toInt(o[0])
			option = &domain.Option{ID: id, Text: strings.Trim(o[1], " ")}
			question.Options[id] = option
		}
	}

	return questionaire
}

func toInt(s string) int {
	i, err := strconv.Atoi(strings.Trim(s, " "))
	if err != nil {
		log.Fatal(err)
	}
	return i
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
