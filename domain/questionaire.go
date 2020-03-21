package domain

type Questionaire struct {
	Questions map[int]Question // Questions maps Questions by their IDs
}

type consumer interface {
	Consume(i interface{}, o interface{}) error
}

func NewQuestionaire(data []byte, c consumer) (Questionaire, error) {
	q := Questionaire{}
	q.Questions = make(map[int]Question)
	err := c.Consume(data, &q.Questions)
	return q, err
}
