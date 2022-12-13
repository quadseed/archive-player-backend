package handler

import "net/http"

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World, Echo World"))
}
