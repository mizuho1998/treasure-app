package service

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
)

type Keyboard struct {
	db *sqlx.DB
}

func NewKeyboardService(db *sqlx.DB) *Keyboard {
	return &Keyboard{db}
}

func (a *Keyboard) Serch(requestSerchKeyboard *model.RequestSerchKeyboard) (*[]model.KeyboardID, error) {
	fmt.Println("service/keyboard Serch()")
	qid := requestSerchKeyboard.QID
	serchKeyboard := &model.SerchKeyboard{}
	serchKeyboard.IDs = requestSerchKeyboard.IDs
	serchKeyboard.Answer = requestSerchKeyboard.Answer

	keyboards := &[]model.KeyboardID{}
	var err error

	fmt.Println("#############")
	fmt.Println(qid)
	fmt.Println("#############")

	switch qid {
	case 1:
		keyboards, err = repository.SerchKeyboardBySplit(a.db, serchKeyboard)
	case 2:
		keyboards, err = repository.SerchKeyboardByLed(a.db, serchKeyboard)
	case 3, 6, 7:
		keyboards, err = repository.SerchKeyboardByMatrix(a.db, serchKeyboard)
	case 4:
		if serchKeyboard.Answer == "0" {
			keyboards, err = repository.SerchKeyboardByKeyNumSmall(a.db, serchKeyboard)
		} else {
			keyboards, err = repository.SerchKeyboardByKeyNumLearge(a.db, serchKeyboard)
		}
	case 5:
		keyboards, err = repository.SerchKeyboardByProfile(a.db, serchKeyboard)
	}
	fmt.Println(keyboards)

	if err != nil && err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return keyboards, nil
}
