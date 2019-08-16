package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
)

type Keyboard struct {
	db *sqlx.DB
}

func NewKeyboard(db *sqlx.DB) *Keyboard {
	return &Keyboard{db}
}

func (a *Keyboard) Update(id int64, newKeyboard *model.Keyboard) error {
	_, err := repository.FindKeyboard(a.db, id)
	if err != nil {
		return errors.Wrap(err, "failed find keyboard")
	}

	// if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
	// 	_, err := repository.UpdateKeyboard(tx, id, newKeyboard)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if err := tx.Commit(); err != nil {
	// 		return err
	// 	}
	// 	return err
	// }); err != nil {
	// 	return errors.Wrap(err, "failed keyboard update transaction")
	// }

	return nil
}
