package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rafaelmgr12/go-server-techniques/custom_api/users"
)

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var request AddUserRequest
	err = json.Unmarshal(data, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users.AddUser(request.Name, request.Age)
	w.WriteHeader(http.StatusOK)
}
