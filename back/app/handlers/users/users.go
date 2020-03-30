package users

import (
	users "blog/back/app/db/users"
	"blog/back/app/libs/reply"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/mitchellh/mapstructure"
)

// GetUser model
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var usersModel []User

	u, err := users.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	mapstructure.Decode(u, &usersModel)

	reply.JSON(w, http.StatusOK, &usersModel)

	return
}

func Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var (
		id        int
		err       error
		userModel User
	)

	if id, err = strconv.Atoi(p.ByName("id")); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := users.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	mapstructure.Decode(u, &userModel)

	reply.JSON(w, http.StatusOK, &userModel)

	return
}
