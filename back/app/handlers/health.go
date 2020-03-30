package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Health is used for checking if server is running
func Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Up and running!\n"))
}
