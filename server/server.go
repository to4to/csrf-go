package server

import (
	"log"
	"net/http"

	"github.com/to4to/csrf-go/server/middleware"
)

func startServer(hostname, port string) error {

	host := hostname + ":" + port

	log.Printf("Listening on : %s", host)

	handler := middleware.NewHandler()
	http.Handle("/", handler)

	return http.ListenAndServe(host, nil)

}
