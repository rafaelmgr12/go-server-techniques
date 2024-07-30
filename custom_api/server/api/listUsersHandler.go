package api

import (
	"io"
	"net/http"
)

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://fakerapi.it/api/v1/persons")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if res.StatusCode != 200 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resData, _ := io.ReadAll(res.Body)
	w.Write(resData)

}
