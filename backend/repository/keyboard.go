package repository

import (
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func AllKeyboard(db *sqlx.DB) (*[]model.KeyboardID, error) {
	keyboardList := []model.KeyboardID{}

	rows, err := db.Queryx(`SELECT id from keyboards`)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		k := model.KeyboardID{}
		err = rows.StructScan(&k)
		keyboardList = append(keyboardList, k)
	}

	return &keyboardList, nil
}

func FindKeyboard(db *sqlx.DB, id int64) (*model.Keyboard, error) {
	a := model.Keyboard{}
	if err := db.Get(&a, `
SELECT id, name, creater_name, price, url, image_url FROM keyboards WHERE id = ?
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

func SerchKeyboard(db *sqlx.DB, rsk *model.SerchKeyboard) (*[]model.KeyboardID, error) {
	keyboardList := []model.KeyboardID{}

	args := make([]interface{}, len(rsk.IDs))
	for i, kid := range rsk.IDs {
		args[i] = kid.ID
	}

	stmt := `SELECT id from keyboards where id in (?` + strings.Repeat(",?", len(rsk.IDs)-1) + `)`
	rows, err := db.Queryx(stmt, args...)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		k := model.KeyboardID{}
		err = rows.StructScan(&k)
		keyboardList = append(keyboardList, k)
	}

	return &keyboardList, nil
}

func SerchKeyboardBySplit(db *sqlx.DB, sk *model.SerchKeyboard) (*[]model.KeyboardID, error) {
	keyboardList := []model.KeyboardID{}

	args := make([]interface{}, len(sk.IDs)+1)
	for i, kid := range sk.IDs {
		args[i] = kid.ID
	}
	args[len(sk.IDs)] = sk.Answer

	stmt := `SELECT id from keyboards where id in (?` + strings.Repeat(",?", len(sk.IDs)-1) + `) AND split = ?`
	rows, err := db.Queryx(stmt, args...)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		k := model.KeyboardID{}
		err = rows.StructScan(&k)
		keyboardList = append(keyboardList, k)
	}

	return &keyboardList, nil
}

func SerchKeyboardByLed(db *sqlx.DB, sk *model.SerchKeyboard) (*[]model.KeyboardID, error) {
	keyboardList := []model.KeyboardID{}

	args := make([]interface{}, len(sk.IDs)+1)
	for i, kid := range sk.IDs {
		args[i] = kid.ID
	}
	args[len(sk.IDs)] = sk.Answer

	stmt := `SELECT id from keyboards where id in (?` + strings.Repeat(",?", len(sk.IDs)-1) + `) AND led = ?`
	rows, err := db.Queryx(stmt, args...)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		k := model.KeyboardID{}
		err = rows.StructScan(&k)
		keyboardList = append(keyboardList, k)
	}

	return &keyboardList, nil
}

func SerchKeyboardByKeyNumLearge(db *sqlx.DB, sk *model.SerchKeyboard) (*[]model.KeyboardID, error) {
	keyboardList := []model.KeyboardID{}

	args := make([]interface{}, len(sk.IDs))
	for i, kid := range sk.IDs {
		args[i] = kid.ID
	}

	stmt := `SELECT id from keyboards where id in (?` + strings.Repeat(",?", len(sk.IDs)-1) + `) AND key_num >= 65`
	rows, err := db.Queryx(stmt, args...)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		k := model.KeyboardID{}
		err = rows.StructScan(&k)
		keyboardList = append(keyboardList, k)
	}

	return &keyboardList, nil
}

func SerchKeyboardByKeyNumSmall(db *sqlx.DB, sk *model.SerchKeyboard) (*[]model.KeyboardID, error) {
	keyboardList := []model.KeyboardID{}

	args := make([]interface{}, len(sk.IDs))
	for i, kid := range sk.IDs {
		args[i] = kid.ID
	}

	stmt := `SELECT id from keyboards where id in (?` + strings.Repeat(",?", len(sk.IDs)-1) + `) AND key_num < 65`
	rows, err := db.Queryx(stmt, args...)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		k := model.KeyboardID{}
		err = rows.StructScan(&k)
		keyboardList = append(keyboardList, k)
	}

	return &keyboardList, nil
}

func SerchKeyboardByMatrix(db *sqlx.DB, sk *model.SerchKeyboard) (*[]model.KeyboardID, error) {
	keyboardList := []model.KeyboardID{}

	args := make([]interface{}, len(sk.IDs)+1)
	for i, kid := range sk.IDs {
		args[i] = kid.ID
	}
	args[len(sk.IDs)] = sk.Answer

	stmt := `SELECT id from keyboards where id in (?` + strings.Repeat(",?", len(sk.IDs)-1) + `) AND matrix = ?`
	rows, err := db.Queryx(stmt, args...)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		k := model.KeyboardID{}
		err = rows.StructScan(&k)
		keyboardList = append(keyboardList, k)
	}

	return &keyboardList, nil
}

func SerchKeyboardByProfile(db *sqlx.DB, sk *model.SerchKeyboard) (*[]model.KeyboardID, error) {
	keyboardList := []model.KeyboardID{}

	args := make([]interface{}, len(sk.IDs)+1)
	for i, kid := range sk.IDs {
		args[i] = kid.ID
	}
	args[len(sk.IDs)] = sk.Answer

	stmt := `SELECT id from keyboards where id in (?` + strings.Repeat(",?", len(sk.IDs)-1) + `) AND key_profile = ?`
	rows, err := db.Queryx(stmt, args...)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		k := model.KeyboardID{}
		err = rows.StructScan(&k)
		keyboardList = append(keyboardList, k)
	}

	return &keyboardList, nil
}

// split, led, key_num, matrix, key_profile, color
