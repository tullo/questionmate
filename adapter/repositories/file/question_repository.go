package file

import (
	"fmt"
	"io/ioutil"
	"os"

	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/rwirdemann/questionmate/domain"
)

type QuestionRepository struct {
	questions []domain.Question
}

func NewQuestionRepository(file string) QuestionRepository {
	fn := fmt.Sprintf("%s/src/github.com/rwirdemann/questionmate/config/%s", os.Getenv("GOPATH"), file)
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	r := QuestionRepository{}
	r.questions = LoadQuestions(data)
	return r
}

func (q QuestionRepository) GetQuestions() []domain.Question {
	return q.questions
}

// todo parsing des byte streams in parser auslagen
func LoadQuestions(data []byte) []domain.Question {
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

func isTarget(s string) bool {
	match, _ := regexp.MatchString("(^ {2}- [a-z]+): \\d+", s)
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
