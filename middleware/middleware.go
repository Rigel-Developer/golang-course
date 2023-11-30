package middleware

import (
	"fmt"
	"net/http"
)

func MiMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Antes de ejecutar el handler")
		next.ServeHTTP(w, r)
		fmt.Println("Despues de ejecutar el handler")
	}
}
