package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func AllQuestion(db *sqlx.DB) ([]model.Question, error) {
	a := make([]model.Question, 0)
	if err := db.Select(&a, `SELECT id, title, body FROM question`); err != nil {
		return nil, err
	}
	return a, nil
}

func FindQuestion(db *sqlx.DB, id int64) (*model.Question, error) {
	a := model.Question{}
	if err := db.Get(&a, `
SELECT id FROM questions WHERE id = ?
`, id); err != nil {
		return nil, err
	}
	return &a, nil
}
