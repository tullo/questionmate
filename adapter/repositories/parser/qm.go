package parser

import (
	"github.com/rwirdemann/questionmate/domain"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type QMParser struct {
}

func (p QMParser) Suffix() string {
	return "qm"
}

func (p QMParser) ParseTargets(data []byte) map[string]string {
	lines := strings.Split(string(data), "\n")
	targets := make(map[string]string)
	for _, l := range lines {
		t := strings.Split(l, ":")
		targets[strings.Trim(t[0], " ")] = strings.Trim(t[1], " ")
	}
	return targets
}

func (p QMParser) ParseDescriptions(data []byte) map[int]string {
	lines := strings.Split(string(data), "\n")
	descriptions := make(map[int]string)
	for _, l := range lines {
		d := strings.Split(l, ":")
		questionID := toInt(d[0])
		descriptions[questionID] = strings.Trim(d[1], " ")
	}
	return descriptions
}

func (p QMParser) ParseQuestions(data []byte) []domain.Question {
	lines := strings.Split(string(data), "\n")
	var questions []domain.Question
	var option *domain.Option

	for _, l := range lines {
		if isQuestion(l) {
			q := strings.Split(l, ":")
			id := toInt(q[0])
			questions = append(questions, domain.NewQuestion(id, strings.Trim(q[1], " ")))
		}

		if len(questions) > 0 && isType(l) {
			t := strings.Split(l, ":")
			questions[len(questions)-1].Type = strings.Trim(t[1], " ")
		}
		if len(questions) > 0 && isDesc(l) {
			t := strings.Split(l, ":")
			questions[len(questions)-1].Desc = strings.Trim(t[1], " ")
		}

		if len(questions) > 0 && isDependency(l) {
			d := strings.Split(l, "=>")
			questionID := toInt(d[0])
			optionID := toInt(d[1])
			questions[len(questions)-1].Dependencies[questionID] = optionID
		}

		if len(questions) > 0 && isOption(l) {
			o := strings.Split(l, ":")
			id := toInt(o[0])
			option = domain.NewOption(id, strings.Trim(o[1], " "))
			questions[len(questions)-1].Options = append(questions[len(questions)-1].Options, option)
		}

		if option != nil && isTarget(l) {
			t := strings.Split(l, ":")
			target := strings.Trim(t[0], " ")
			value := toInt(t[1])
			option.Targets[target[2:]] = domain.Score{Value: value}
		}
	}

	return questions
}

func isDesc(s string) bool {
	return strings.HasPrefix(s, "desc:")
}

func isTarget(s string) bool {
	match, _ := regexp.MatchString("(^ {2}- [a-zäÄöÖüÜß]+): \\d+", s)
	return match
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

func isDependency(s string) bool {
	match, _ := regexp.MatchString("(^ {2}\\d+) => \\d+", s)
	return match
}
