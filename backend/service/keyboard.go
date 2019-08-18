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
	qid := requestSerchKeyboard.QID
	serchKeyboard := &model.SerchKeyboard{}
	serchKeyboard.IDs = requestSerchKeyboard.IDs
	serchKeyboard.Answer = requestSerchKeyboard.Answer

	fmt.Println("!!!!!!!!!!!!!")
	fmt.Println(requestSerchKeyboard)
	fmt.Println(requestSerchKeyboard.Answer)
	fmt.Println(serchKeyboard.Answer)
	fmt.Println(serchKeyboard)

	keyboards := &[]model.KeyboardID{}
	var err error

	switch qid {
	case 1:
		keyboards, err = repository.SerchKeyboardBySplit(a.db, serchKeyboard)
		fmt.Println("split")
		fmt.Println(keyboards)
	case 2:
		keyboards, err = repository.SerchKeyboardByMatrix(a.db, serchKeyboard)
		fmt.Println("matrix")
		fmt.Println(keyboards)
	case 3:
		if serchKeyboard.Answer == "0" {
			keyboards, err = repository.SerchKeyboardByKeyNumSmall(a.db, serchKeyboard)
		} else {
			keyboards, err = repository.SerchKeyboardByKeyNumLearge(a.db, serchKeyboard)
		}
		fmt.Println("key_num")
		fmt.Println(keyboards)
	case 4:
		keyboards, err = repository.SerchKeyboardByProfile(a.db, serchKeyboard)
		fmt.Println("profile")
		fmt.Println(keyboards)
	case 5:
		keyboards, err = repository.SerchKeyboardByLed(a.db, serchKeyboard)
		fmt.Println("led")
		fmt.Println(keyboards)
	}

	if err != nil && err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return keyboards, nil
}
