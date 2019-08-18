package model

type Question struct {
	ID         int64   `db:"id" json:"id"`
	Question   string  `db:"question" json:"question"`
	Answer1    string  `db:"answer1" json:"answer1"`
	Answer2    string  `db:"answer2" json:"answer2"`
	Answer3    *string `db:"answer3" json:"answer3"`
	Answer1Val string  `db:"answer1_val" json:"answer1_val"`
	Answer2Val string  `db:"answer2_val" json:"answer2_val"`
	Answer3Val *string `db:"answer3_val" json:"answer3_val"`
}
