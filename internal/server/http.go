package server

import (
	"log"
	"net/http"
	"text/template"

	"github.com/SimilarEgs/L0-orders/pkg/cache"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("./static/templates/*.html"))
}

func OrderByIdHandler(w http.ResponseWriter, r *http.Request) {

	orderId := r.URL.Query().Get("id")
	order, ok := cache.AppCache.Get(orderId)

	if !ok {
		http.NotFound(w, r)
		return
	}

	err := tmpl.ExecuteTemplate(w, "orders.html", order)
	if err != nil {
		log.Fatal(err)
	}
}
