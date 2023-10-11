package mid

import (
	"fmt"
	"net/http"
)

func ContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.Header().Set("Basic", "true")

		next.ServeHTTP(w, r)
	})
}

func ValidRequest(r *http.Request) bool {
	fmt.Println("Página deu PAU: ", r.Method)
	return true
}
