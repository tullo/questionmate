package domain

import (
	"fmt"
	"sort"
)

type Question struct {
	ID      int
	Text    string
	Type    string
	Options []*Option
}

func NewQuestion(id int, text string) *Question {
	q := Question{ID: id, Text: text}
	return &q
}

func (q Question) GetOption(id int) *Option {
	for _, option := range q.Options {
		if option.ID == id {
			return option
		}
	}
	return nil
}

type Option struct {
	ID      int
	Text    string
	Targets map[string]Score
}

func NewOption(id int, text string) *Option {
	o := Option{ID: id, Text: text}
	o.Targets = make(map[string]Score)
	return &o
}

type Score struct {
	Value int
}

type Questionaire struct {
	Questions map[int]*Question // Questions maps Questions by their IDs
}

func NewQuestionnaire() Questionaire {
	q := Questionaire{}
	q.Questions = make(map[int]*Question)
	return q
}

// Unanswered returns an sorted array of unanswered question ids according to
// the given answers.
func (q Questionaire) Unanswered(answers []Answer) []int {
	var unanswered []int
	for id := range q.Questions {
		if !contains(id, answers) {
			unanswered = append(unanswered, id)
		}
	}
	sort.Ints(unanswered)
	return unanswered
}

func contains(id int, answers []Answer) bool {
	for _, a := range answers {
		if a.ID == id {
			return true
		}
	}
	return false
}

func (q Questionaire) String() string {
	var s string
	for _, question := range q.Questions {
		s = fmt.Sprintf("%sQuestion: %s\n", s, question.Text)
		for _, option := range question.Options {
			s = fmt.Sprintf("%s- %s\n", s, option.Text)
		}
	}
	return s
}

type Answer struct {
	ID int
}
