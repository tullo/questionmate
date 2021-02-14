package file

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/tullo/questionmate/adapter/repositories/parser"
	"github.com/tullo/questionmate/domain"
)

type QuestionnaireRepository struct {
	questionnaire domain.Questionnaire
	Questions     []domain.Question
	Descriptions  map[int]string
	Targets       map[string]string
	Ratings       map[string][]domain.Rating
}

func NewQuestionnaireRepository(filename string, p parser.Parser) QuestionnaireRepository {
	var q domain.Questionnaire
	var qs []domain.Question
	var rm map[string][]domain.Rating
	var dm map[int]string
	var tm map[string]string

	if bytes, ok := readFile(filename + "." + p.Suffix()); ok {
		qs = p.ParseQuestions(bytes)
		q = p.ParseQuestionnaire(bytes)
	}

	return QuestionnaireRepository{questionnaire: q, Questions: qs,
		Ratings: rm, Descriptions: dm, Targets: tm}
}

func readFile(file string) ([]byte, bool) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return nil, false
	}
	return data, true
}

func (qr QuestionnaireRepository) GetQuestions() []domain.Question {
	return qr.Questions
}

func (qr QuestionnaireRepository) GetRatings() map[string][]domain.Rating {
	return qr.Ratings
}

func (qr QuestionnaireRepository) GetDescriptions() map[int]string {
	return qr.Descriptions
}

func (qr QuestionnaireRepository) GetQuestionnaire() domain.Questionnaire {
	return qr.questionnaire
}

func (qr QuestionnaireRepository) String() string {
	var s string
	for _, q := range qr.Questions {
		s = fmt.Sprintf("%s---------------------------------------------------------------------------------------\n", s)
		s = fmt.Sprintf("%sFrage: %s\n", s, q.Text)
		if desc, ok := qr.Descriptions[q.ID]; ok {
			s = fmt.Sprintf("%sBeschreibung: %s\n", s, desc)
		}
		for _, o := range q.Options {
			s = fmt.Sprintf("%s- %s\n", s, o.Text)
			for k, v := range o.Targets {
				s = fmt.Sprintf("%s  - %s: %d\n", s, k, v.Value)
			}
		}
	}
	return s
}
