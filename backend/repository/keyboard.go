package repository

import (
	"fmt"
	"strconv"
	"strings"

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

func intToStringArray(arr []int) []string {
	f := make([]string, len(arr))
	for n := range arr {
		f[n] = strconv.Itoa(arr[n])
	}
	return f
}

func SerchKeyboard(db *sqlx.DB, rsk *model.RequestSerchKeyboard) (*[]model.Keyboard, error) {
	keyboardList := []model.Keyboard{}

	args := make([]interface{}, len(rsk.IDs))
	for i, id := range rsk.IDs {
		args[i] = id
	}

	stmt := `SELECT id from keyboards where id in (?` + strings.Repeat(",?", len(rsk.IDs)-1) + `)`
	rows, err := db.Queryx(stmt, args...)

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
