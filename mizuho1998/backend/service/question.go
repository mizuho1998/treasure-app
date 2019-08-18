package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
)

type Question struct {
	db *sqlx.DB
}

func NewQuestion(db *sqlx.DB) *Question {
	return &Question{db}
}

func (a *Question) Update(id int64, newQuestion *model.Question) error {
	_, err := repository.FindQuestion(a.db, id)
	if err != nil {
		return errors.Wrap(err, "failed find question")
	}

	// if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
	// 	_, err := repository.UpdateQuestion(tx, id, newQuestion)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if err := tx.Commit(); err != nil {
	// 		return err
	// 	}
	// 	return err
	// }); err != nil {
	// 	return errors.Wrap(err, "failed question update transaction")
	// }
	return nil
}
