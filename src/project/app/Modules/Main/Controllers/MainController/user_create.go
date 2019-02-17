package MainController

import (
	"net/http"
	"log"
	"project/lib/render"
	"project/lib/api"
	"project/app/Modules/Main/Models/UserModel"
	"project/lib/response"
	"io/ioutil"
	"project/lib/errors"
	"project/lib/header"
	"encoding/json"
	"time"
)

type OutUserCreate struct {
	header.Header
	FormNameError   string `json:"formNameError"`
	FormAvatarError string `json:"formAvatarError"`
	Id string `json:"id"`
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		out := new(OutUserCreate)
		defer response.SendResponse(w, out, "OutUserCreate")

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("###ERR: MainController->UserEdit->ioutil.ReadAll(r.Body):", err)
			out.SetError(errors.ERROR_INTERNAL)
			return
		}

		user := new(UserModel.User)
		err = json.Unmarshal(data, &user)
		if err != nil {
			log.Println("###ERR: MainController->UserCreate->json.Unmarshal(data, &user)1:", err)
			out.SetError(errors.ERROR_INTERNAL)
			return
		}

		user.CreatedAt = time.Now().Format("2006-01-02T15:04:05Z")

		formErr := false
		if user.Name == "" {
			out.FormNameError = "Поле Name не может быть пустым"
			formErr = true
		}

		if user.Avatar == "" {
			out.FormAvatarError = "Поле Avatar не может быть пустым"
			formErr = true
		}

		if formErr {
			return
		}

		var code int
		code, data = api.NewApiRequest("/user/", http.MethodPost, user)
		if code != http.StatusCreated {
			log.Println("###ERR: MainController->UserCreate: code api:", code)
			out.SetError(errors.ERROR_INTERNAL)
			return
		}

		err = json.Unmarshal(data, &user)
		if err != nil {
			log.Println("###ERR: MainController->UserCreate->json.Unmarshal(data, &user)2:", err)
			out.SetError(errors.ERROR_INTERNAL)
			return
		}

		out.Id = user.Id

		return
	}

	err := render.Render(w)
	if err != nil {
		log.Println("###ERR: MainController->UserCreate->render.Render(w):", err)
		return
	}
}
