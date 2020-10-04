package parser

import (
	"github.com/rwirdemann/questionmate/domain"
	"gopkg.in/yaml.v2"
	"log"
)

type question struct {
	Id      int      `yaml:"id"`
	Text    string   `yaml:"text"`
	Options []option `yaml:"options"`
}

type option struct {
	Value   int      `yaml:"value"`
	Text    string   `yaml:"text"`
	Targets []target `yaml:"targets"`
}

type target struct {
	Name  string `yaml:"name"`
	Score int    `yaml:"score"`
}

type t struct {
	Questions []question `yaml:",flow"`
}

type YAMLParser struct {
}

func (Y YAMLParser) Suffix() string {
	return "yaml"
}

func (Y YAMLParser) ParseQuestions(data []byte) []domain.Question {
	m := t{}
	err := yaml.Unmarshal(data, &m)
	if err != nil {
		log.Fatal(err)
	}

	var questions []domain.Question
	for _, v := range m.Questions {
		q := domain.NewQuestion(v.Id, v.Text)
		for _, o := range v.Options {
			option := domain.NewOption(o.Value, o.Text)
			for _, t := range o.Targets {
				option.Targets[t.Name] = domain.Score{Value: t.Score}
			}
			q.Options = append(q.Options, option)
		}
		questions = append(questions, q)

	}
	return questions
}

func (Y YAMLParser) ParseDescriptions([]byte) map[int]string {
	panic("implement me")
}

func (Y YAMLParser) ParseTargets([]byte) map[string]string {
	panic("implement me")
}
