package usecase

import "github.com/rwirdemann/questionmate/domain"

type GetEvaluation struct {
	QuestionRepository questionRepository
}

func (e GetEvaluation) Process(answers []domain.Answer) domain.Evaluation {
	questions := e.QuestionRepository.GetQuestions()
	evaluation := domain.Evaluation{}
	for _, answer := range answers {
		if q, ok := questionByID(questions, answer.QuestionID); ok {
			if o, ok := q.GetOption(answer.Value); ok {
				for test, score := range o.Targets {
					if t, ok := evaluation.GetTarget(test); ok {
						t.Score += score.Value
					} else {
						evaluation.Targets = append(evaluation.Targets, &domain.Target{
							Text:  test,
							Score: score.Value,
						})
					}
				}
			}
		}
	}

	return evaluation
}

func questionByID(questions []domain.Question, id int) (domain.Question, bool) {
	for _, question := range questions {
		if question.ID == id {
			return question, true
		}
	}
	return domain.Question{}, false
}
