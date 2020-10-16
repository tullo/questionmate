package parser

import (
	"log"

	"github.com/rwirdemann/questionmate/domain"
	"gopkg.in/yaml.v2"
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
	Description string     `yaml:"description"`
	Questions   []question `yaml:",flow"`
}

type value struct {
	Min         int    `yaml:"min"`
	Max         int    `yaml:"max"`
	Description string `yaml:"description"`
}

type ratings struct {
	Target string  `yaml:"target"`
	Values []value `yaml:",flow"`
}

type a struct {
	Ratings []ratings `yaml:",flow"`
}

type questionnaire struct {
	Abstract string `yaml:"abstract"`
}

type YAMLParser struct {
}

func (Y YAMLParser) ParseRatings(data []byte) map[string][]domain.Rating {
	m := a{}
	err := yaml.Unmarshal(data, &m)
	if err != nil {
		log.Fatal(err)
	}

	ranges := make(map[string][]domain.Rating)
	for _, v := range m.Ratings {
		assessmentRanges := ranges[v.Target]
		for _, r := range v.Values {
			assessmentRanges = append(assessmentRanges, domain.Rating{
				Target:      v.Target,
				Min:         r.Min,
				Max:         r.Max,
				Description: r.Description,
			})
		}
		ranges[v.Target] = assessmentRanges
	}
	return ranges
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

func (Y YAMLParser) ParseQuestionnaire(data []byte) domain.Questionnaire {
	q := questionnaire{}
	err := yaml.Unmarshal(data, &q)
	if err != nil {
		log.Fatal(err)
	}

	return domain.Questionnaire{
		Abstract: q.Abstract,
	}
}
