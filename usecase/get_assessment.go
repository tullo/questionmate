package usecase

import "github.com/rwirdemann/questionmate/domain"

type Assessment struct {
	QuestionRepository questionRepository
}

func (e Assessment) GetAssessment(answers []domain.Answer) domain.Assessment {
	questions := e.QuestionRepository.GetQuestions()
	assessment := domain.Assessment{}
	for _, answer := range answers {
		if q, ok := questionByID(questions, answer.QuestionID); ok {
			if o, ok := q.GetOption(answer.Value); ok {
				for test, score := range o.Targets {
					if t, ok := assessment.GetTarget(test); ok {
						t.Score += score.Value
					} else {
						assessment.Targets = append(assessment.Targets, &domain.Target{
							Text:  test,
							Score: score.Value,
						})
					}
				}
			}
		}
	}

	ratings := e.QuestionRepository.GetRatings()
	for _, t := range assessment.Targets {
		if values, ok := ratings[t.Text]; ok {
			for _, v := range values {
				if t.Score >= v.Min && t.Score <= v.Max {
					t.Rating = v.Description
					break
				}
			}
		}
	}

	return assessment
}

func questionByID(questions []domain.Question, id int) (domain.Question, bool) {
	for _, question := range questions {
		if question.ID == id {
			return question, true
		}
	}
	return domain.Question{}, false
}
