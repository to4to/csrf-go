package middleware

import (
	"net/http"

	"github.com/justinas/alice"
)

func NewHandler() http.Handler {

	return alice.New(recoverHandler, authHandler).ThenFunc(logicHandler)

}

func logicHandler(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/restricted":
	case "/login":
		switch r.Method {
		case "GET":
		case "POST":
		default:
		}
	}
}
