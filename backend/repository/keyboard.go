package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func AllKeyboard(db *sqlx.DB) ([]model.Keyboard, error) {
	a := make([]model.Keyboard, 0)
	if err := db.Select(&a, `SELECT id, title, body FROM keyboard`); err != nil {
		return nil, err
	}
	return a, nil
}

func FindKeyboard(db *sqlx.DB, id int64) (*model.Keyboard, error) {
	a := model.Keyboard{}
	if err := db.Get(&a, `
SELECT id, name, creater_name, url FROM keyboards WHERE id = ?
`, id); err != nil {
		return nil, err
	}
	return &a, nil
}

func SerchKeyboard(db *sqlx.DB, rsk *model.RequestSerchKeyboard) (*[]model.Keyboard, error) {
	// 	a := model.Keyboard{}
	// 	if err := db.Get(&a, `
	// SELECT id FROM keyboards
	// `); err != nil {
	// 		return nil, err
	// 	}
	keyboardList := []model.Keyboard{}

	fmt.Println(rsk)

	rows, err := db.Queryx(`
	SELECT id FROM keyboards where is_split = ? AND color = ? AND key_num = ? AND matrix = ? AND key_profile = ?
	`, rsk.IsSplit, rsk.Color, rsk.KeyNum, rsk.Matrix, rsk.KeyProfile)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		k := model.Keyboard{}
		err = rows.StructScan(&k)
		keyboardList = append(keyboardList, k)
	}
	fmt.Println(keyboardList)

	return &keyboardList, nil
}
