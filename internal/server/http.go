package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ServeRoutes() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/orders/{id}", OrderByIdHandler)

	return r
}

func OrderByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintf(w, "id of the order - %s", id)
}
