package MainController

import (
	"net/http"
	"project/lib/api"
	"fmt"
	"encoding/json"
	"project/app/Modules/Main/Models/UserModel"
	"project/lib/render"
	"project/lib/header"
	"project/lib/response"
	"project/lib/errors"
)

type OutUsersList struct {
	header.Header
	List []*UserModel.User `json:"list"`
}

func UsersList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		out := new(OutUsersList)
		defer response.SendResponse(w, out, "OutUsersList")

		code, data := api.NewApiRequest("/user", http.MethodGet, nil)
		if code != http.StatusOK {
			fmt.Println("###ERR: MainController->UsersList: code api:", code)
			out.SetError(errors.ERROR_INTERNAL)
			return
		}

		var usersList []*UserModel.User

		err := json.Unmarshal(data, &usersList)
		if err != nil {
			fmt.Println("###ERR: MainController->UsersList->json.Unmarshal(data, &usersList):", err)
			out.SetError(errors.ERROR_INTERNAL)
			return
		}

		out.List = usersList
		return
	}

	err := render.Render(w)
	if err != nil {
		fmt.Println("###ERR: MainController->UsersList->render.Render(w):", err)
		return
	}
}
