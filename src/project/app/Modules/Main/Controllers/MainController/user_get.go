package MainController

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"project/lib/api"
	"project/lib/render"
	"project/lib/header"
	"project/lib/response"
	"project/lib/errors"
	"encoding/json"
	"project/app/Modules/Main/Models/UserModel"
)

type OutUserGet struct {
	header.Header
	User *UserModel.User `json:"user"`
}

func UserGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		out := new(OutUserGet)
		defer response.SendResponse(w, out, "OutUserGet")

		vars := mux.Vars(r)
		id := vars["id"]

		code, data := api.NewApiRequest("/user/" + id, http.MethodGet, nil)
		if code != http.StatusOK {
			log.Println("###ERR: MainController->UserGet: code api:", code)
			if code == errors.ERROR_NOT_FOUND {
				out.SetError(errors.ERROR_NOT_FOUND)
			} else {
				out.SetError(errors.ERROR_INTERNAL)
			}
			return
		}

		user := new(UserModel.User)

		err := json.Unmarshal(data, &user)
		if err != nil {
			log.Println("###ERR: MainController->UserGet->json.Unmarshal(data, &user):", code)
			out.SetError(errors.ERROR_INTERNAL)
			return
		}

		out.User = user
		return
	}

	err := render.Render(w)
	if err != nil {
		log.Println("###ERR: MainController->UserGet->render.Render(w):", err)
		return
	}
}
