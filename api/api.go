package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const portNumber = ":80"

func Init() {
	r := &router{mux.NewRouter()}
	r.Use(mux.CORSMethodMiddleware(r.Router), LoggingMiddleware(r))

	addUserHandlers(r)

	r.handleGet("/books/{title}/page/{page}", getBook)

	http.ListenAndServe(portNumber, r)
}

func CreateUrl(base string, args ...string) string {
	result := base
	for _, seg := range args {
		result += "/" + seg
	}
	return result
}

func getBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title, page := vars["title"], vars["page"]
	fmt.Fprintf(w, "You have requested the book: %s on page %s\n", title, page)
}
