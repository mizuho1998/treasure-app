package model

type Question struct {
	ID       int64  `db:"id" json:"id"`
	Question string `db:"question" json:"question"`
}
