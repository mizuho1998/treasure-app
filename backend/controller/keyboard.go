package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/voyagegroup/treasure-app/httputil"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
	"github.com/voyagegroup/treasure-app/service"
)

type Keyboard struct {
	db *sqlx.DB
}

func NewKeyboard(db *sqlx.DB) *Keyboard {
	return &Keyboard{db: db}
}

func (a *Keyboard) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	keyboards, err := repository.AllKeyboard(a.db)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, keyboards, nil
}

func (a *Keyboard) AllKeyboard(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {

	keyboard, err := repository.AllKeyboard(a.db)
	if err != nil && err == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	fmt.Println("controller/keyboard AllKeyboard()")
	fmt.Println(keyboard)

	return http.StatusCreated, keyboard, nil
}

func (a *Keyboard) FindKeyboard(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	keyboard, err := repository.FindKeyboard(a.db, aid)
	if err != nil && err == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, keyboard, nil
}

func (a *Keyboard) Serch(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	fmt.Println("keyboard controller Serch")

	serchKeyboard := &model.RequestSerchKeyboard{}
	if err := json.NewDecoder(r.Body).Decode(&serchKeyboard); err != nil {
		return http.StatusBadRequest, nil, err
	}

	// serchKeyboard := &model.RequestSerchKeyboard{
	// 	IDs:    reqParam.IDs,
	// 	QID:    reqParam.QID,
	// 	Answer: reqParam.Answer,
	// }
	fmt.Println(serchKeyboard)

	// keyboards, err := repository.SerchKeyboard(a.db, serchKeyboard)
	// if err != nil && err == sql.ErrNoRows {
	// 	return http.StatusNotFound, nil, err
	// } else if err != nil {
	// 	return http.StatusInternalServerError, nil, err
	// }
	keyboardService := service.NewKeyboardService(a.db)
	keyboards, err := keyboardService.Serch(serchKeyboard)
	if err != nil && err == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	fmt.Println("controller Serch() end")
	fmt.Println(keyboards)
	fmt.Println(reflect.TypeOf(keyboards))

	return http.StatusCreated, keyboards, nil
}
