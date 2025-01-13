package handler

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("welcome to student api"))

}
