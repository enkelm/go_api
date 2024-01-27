package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type middleware func(r *router) mux.MiddlewareFunc
type middlewarePreProcess func(w http.ResponseWriter, r *http.Request) any
type middlewarePostProcess func(w http.ResponseWriter, r *http.Request, arg any)

func middlewareBuilder(preFn middlewarePreProcess, postFn middlewarePostProcess) middleware {
	return func(r *router) mux.MiddlewareFunc {
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				res := preFn(w, req)
				next.ServeHTTP(w, req)
				postFn(w, req, res)
			})
		}
	}
}

var LoggingMiddleware = middlewareBuilder(
	func(w http.ResponseWriter, r *http.Request) any {
		start := time.Now()
		return start
	},
	func(w http.ResponseWriter, r *http.Request, arg any) {
		fmt.Println(r.URL.Path, time.Since(arg.(time.Time)))
	},
)
