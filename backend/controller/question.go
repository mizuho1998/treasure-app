package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/voyagegroup/treasure-app/httputil"
	"github.com/voyagegroup/treasure-app/repository"
)

type Question struct {
	db *sqlx.DB
}

func NewQuestion(db *sqlx.DB) *Question {
	return &Question{db: db}
}

func (a *Question) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	questions, err := repository.AllQuestion(a.db)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, questions, nil
}

func (a *Question) Find(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	fmt.Println("controller/question Serch()")
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	qid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	fmt.Printf("id: %d\n", qid)

	question, err := repository.FindQuestion(a.db, qid)
	if err != nil && err == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, question, nil
}
