package domain

import (
	"fmt"
	"sort"
)

type Question struct {
	ID           int       `json:"id"`
	Text         string    `json:"text"`
	Type         string    `json:"type"`
	Options      []*Option `json:"options"`
	Dependencies map[int]int
}

func NewQuestion(id int, text string) *Question {
	q := Question{ID: id, Text: text}
	q.Dependencies = make(map[int]int)
	return &q
}

func (q Question) GetOption(value int) *Option {
	for _, option := range q.Options {
		if option.Value == value {
			return option
		}
	}
	return nil
}

type Option struct {
	Value   int    `json:"value"`
	Text    string `json:"text"`
	Targets map[string]Score
}

func NewOption(value int, text string) *Option {
	o := Option{Value: value, Text: text}
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
		if a.QuestionID == id {
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
	QuestionID int `json:"question_id"`
}
