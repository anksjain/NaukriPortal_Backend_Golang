package middleware

import (
	"fmt"
	"net/http"
)

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do stuff
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Authorization,Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

		fmt.Println("In cores middleware")
		if r.Method == http.MethodOptions {
			fmt.Println("in cores middle for OPTIONS ")
			return
		}
		fmt.Println("out cores middleware")
		h.ServeHTTP(w, r)
	})
}
