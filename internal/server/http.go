package server

import (
	"fmt"
	"net/http"

	"github.com/SimilarEgs/L0-orders/pkg/cache"
	"github.com/gorilla/mux"
)

func ServeRoutes() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/orders/{id}", OrderByIdHandler)

	return r
}

func OrderByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	order, _ := cache.AppCache.Get(id)

	for _, item := range order.Items {
		fmt.Fprintf(w, "items: %v", item)
	}

}
