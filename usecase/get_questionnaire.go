package usecase

import "github.com/rwirdemann/questionmate/domain"

type GetQuestionnaire struct {
	Repositories map[string]questionRepository
}

func NewGetQuestionnaire() GetQuestionnaire {
	return GetQuestionnaire{Repositories: make(map[string]questionRepository)}
}

func (receiver GetQuestionnaire) Process(name string) (domain.Questionnaire, bool) {
	if r, ok := receiver.Repositories[name]; ok {
		return r.GetQuestionnaire(), true
	}
	return domain.Questionnaire{}, false
}
