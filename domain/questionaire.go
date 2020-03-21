package domain

import (
	"gopkg.in/yaml.v2"
)

type Questionaire struct {
	Questions map[int]Question // Questions maps Questions by their IDs
}

func NewQuestionaire(data []byte) (Questionaire, error) {
	q := Questionaire{}
	q.Questions = make(map[int]Question)
	err := yaml.Unmarshal(data, &q.Questions)
	if err != nil {
		return q, err
	}
	return q, nil
}
