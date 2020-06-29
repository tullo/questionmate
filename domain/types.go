package domain

import "fmt"

type Question struct {
	ID      int
	Text    string
	Type    string
	Options map[int]*Option
}

func NewQuestion(id int, text string) *Question {
	q := Question{ID: id, Text: text}
	q.Options = make(map[int]*Option)
	return &q
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
}
