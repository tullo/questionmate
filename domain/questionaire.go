package domain

import (
	"gopkg.in/yaml.v2"
)

type Questionaire struct {
	Questions map[int]Question
}

func NewQuestionaire(data []byte) Questionaire {
	q := Questionaire{}
	q.Questions = make(map[int]Question)
	err := yaml.Unmarshal(data, &q.Questions)
	if err != nil {
		panic(err)
	}
	return q
}
