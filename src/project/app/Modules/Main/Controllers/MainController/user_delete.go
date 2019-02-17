package MainController

import (
	"net/http"
	"project/lib/response"
	"project/lib/header"
	"github.com/gorilla/mux"
	"project/lib/api"
	"fmt"
	"project/lib/errors"
)

type OutUserDelete struct {
	header.Header
	FormStatus bool   `json:"formStatus"`
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		out := new(OutUserDelete)
		defer response.SendResponse(w, out, "OutUserDelete")

		vars := mux.Vars(r)
		id := vars["id"]

		code, _ := api.NewApiRequest("/user/" + id, http.MethodDelete, nil)
		if code != http.StatusOK {
			fmt.Println("###ERR: MainController->UserDelete: code api:", code)
			out.SetError(errors.ERROR_INTERNAL)
			return
		}

		out.FormStatus = true
		return
	}
}
