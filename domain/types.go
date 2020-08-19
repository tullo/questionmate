package domain

import (
	"strconv"
)

type Question struct {
	ID           int       `json:"id"`
	Text         string    `json:"text"`
	Type         string    `json:"type"`
	Desc         string    `json:"desc"`
	Options      []*Option `json:"options"`
	Dependencies map[int]int
}

func NewQuestion(id int, text string) Question {
	q := Question{ID: id, Text: text}
	q.Dependencies = make(map[int]int)
	return q
}

func (q Question) GetOption(value int) (Option, bool) {
	for _, option := range q.Options {
		if option.Value == value {
			return *option, true
		}
	}
	return Option{}, false
}

func (q Question) GetOptionByString(value string) (Option, bool) {
	i, err := strconv.Atoi(value)
	if err != nil {
		return Option{}, false
	}

	return q.GetOption(i)
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

type Answer struct {
	QuestionID int `json:"question_id"`
	Value      int
}

type Target struct {
	Text  string `json:"text"`
	Score int    `json:"score"`
}

type Evaluation struct {
	Targets []*Target `json:"targets"`
}

func (e Evaluation) GetTarget(s string) (*Target, bool) {
	for _, t := range e.Targets {
		if t.Text == s {
			return t, true
		}
	}
	return nil, false
}
