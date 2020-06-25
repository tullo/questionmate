package domain

type Question struct {
	ID      int
	Text    string
	Type    string
	Options map[int]Option
}

type Option struct {
	ID     int
	Text   string
	Scores map[int]Score
}

type Score struct {
	Id    int
	Value int
	Why   string
}

type Target struct {
	ID    int
	Label string
}

type Questionaire struct {
	Targets   map[int]Target
	Questions map[int]*Question // Questions maps Questions by their IDs
}

type Answer struct {
}
