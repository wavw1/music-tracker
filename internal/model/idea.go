package model

type Idea struct {
	ID     int
	UserID int64
	Title  string
	Bpm    int
	Key    string
	Status string
	Tags   []string
}
