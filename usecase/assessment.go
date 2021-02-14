package usecase

import "github.com/tullo/questionmate/domain"

type Assessment struct {
	QR questionRepository
}

func (a Assessment) GetAssessment(as []domain.Answer) domain.Assessment {
	qs := a.QR.GetQuestions()
	assessment := domain.Assessment{}
	for _, answer := range as {
		if q, ok := questionByID(qs, answer.QuestionID); ok {
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

	ratings := a.QR.GetRatings()
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

func questionByID(qs []domain.Question, id int) (domain.Question, bool) {
	for _, q := range qs {
		if q.ID == id {
			return q, true
		}
	}
	return domain.Question{}, false
}
