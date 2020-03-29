package handlers

import (
	"blog/back/app/libs/reply"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Health is used for checking if server is running
func Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		apiResp reply.Respond
	)

	defer reply.JSON(w, http.StatusOK, &apiResp)

}
