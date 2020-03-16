package domain

type Question struct {
	ID      int
	Text    string
	Type    string
	Options map[int]Option
}

type Option struct {
	ID      int
	Text    string
	Targets map[string]int
}
