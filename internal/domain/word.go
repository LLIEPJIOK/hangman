package domain

type Word struct {
	Value    string `json:"word"`
	Attempts int
	Hint     string
}
