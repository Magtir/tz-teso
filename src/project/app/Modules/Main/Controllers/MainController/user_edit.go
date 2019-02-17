package MainController

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"io/ioutil"
	"project/lib/api"
	"project/app/Modules/Main/Models/UserModel"
	"project/lib/header"
	"project/lib/response"
	"project/lib/errors"
	"encoding/json"
)

type OutUserEdit struct {
	header.Header
	FormError string `json:"formError"`
	FormStatus bool   `json:"formStatus"`
}

func UserEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		out := new(OutUserEdit)
		defer response.SendResponse(w, out, "OutUserEdit")

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("###ERR: MainController->UserEdit->ioutil.ReadAll(r.Body):", err)
			out.SetError(errors.ERROR_INTERNAL)
			return
		}

		newName := string(data)

		if newName == "" {
			out.FormError = "Поле Name не может быть пустым"
			return
		}

		vars := mux.Vars(r)
		id := vars["id"]

		code, data := api.NewApiRequest("/user/" + id, http.MethodGet, nil)
		if code != http.StatusOK {
			log.Println("###ERR: MainController->UserEdit: code api 1:", code)
			out.SetError(errors.ERROR_INTERNAL)
			return
		}

		user := new(UserModel.User)

		err = json.Unmarshal(data, &user)
		if err != nil {
			log.Println("###ERR: MainController->UserEdit->json.Unmarshal(data, &user):", code)
			out.SetError(errors.ERROR_INTERNAL)
			return
		}

		user.Name = newName

		code, _ = api.NewApiRequest("/user/"+id, http.MethodPut, user)
		if code != http.StatusOK {
			log.Println("###ERR: MainController->UserEdit: code api 2:", code)
			out.SetError(errors.ERROR_INTERNAL)
			return
		}

		out.FormStatus = true
		return
	}
}
