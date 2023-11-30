package webserver

import (
	"github/rigel-developer/golang-course/middleware"
	"net/http"
)

func MiWebServer() {

	http.HandleFunc("/", middleware.MiMiddleware(home))
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserver/index.html")
}
