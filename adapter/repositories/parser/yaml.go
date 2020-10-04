package parser

import (
	"github.com/rwirdemann/questionmate/domain"
	"gopkg.in/yaml.v2"
	"log"
)

type question struct {
	Id   int    `yaml:"id"`
	Text string `yaml:"text"`
}

type t struct {
	Questions []question `yaml:",flow"`
}

type YAMLParser struct {
}

func (Y YAMLParser) ParseQuestions(data []byte) []domain.Question {
	m := t{}
	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatal(err)
	}

	var questions []domain.Question
	for _, v := range m.Questions {
		q := domain.NewQuestion(v.Id, v.Text)
		questions = append(questions, q)

	}
	return questions
}

func (Y YAMLParser) ParseDescriptions(data []byte) map[int]string {
	panic("implement me")
}

func (Y YAMLParser) ParseTargets(data []byte) map[string]string {
	panic("implement me")
}
