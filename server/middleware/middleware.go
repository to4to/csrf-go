package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/justinas/alice"
	myjwt "github.com/to4to/csrf-go/server/middleware/myJWT"
)

func NewHandler() http.Handler {

	return alice.New(recoverHandler, authHandler).ThenFunc(logicHandler)

}

func recoverHandler(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		defer func() {

			if err := recover(); err != nil {
				log.Panic("Recovered! panic! %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}

		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func authHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/restricted", "/logout", "/deleteUser":

		default:
		}
	}

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
	case "/register":
		switch r.Method {
		case "GET":
		case "POST":
		default:
		}
	case "/logout":
	case "/deleteUser":
	default:
	}
}

func nullifyTokenCookies(w http.ResponseWriter, r http.Request) {

	authCookie := http.Cookie{
		Name:     "AuthToken",
		Value:    "",
		Expires:  time.Now().Add(-1000 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, &authCookie)

	refreshCookie := http.Cookie{
		Name:     "RefreshToken",
		Value:    "",
		Expires:  time.Now().Add(-1000 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, &refreshCookie)

	RefreshCookie, refreshErr := r.Cookie("RefreshToken")

	if refreshErr == http.ErrNoCookie {
		return
	} else if refreshErr != nil {
		log.Panic("panic %+v ", refreshErr)
		http.Error(w, http.StatusText(500), 500)
	}

	myjwt.RevokeRefreshToken(RefreshCookie.Value)
}

func setAuthAndRefreshCookies(w http.ResponseWriter, authTokenString, refreshTokenString string) {
	authCookie := http.Cookie{
		Name:     "AuthToken",
		Value:    authTokenString,
		HttpOnly: true,
	}

	http.SetCookie(w, &authCookie)
	refreshCookie := http.Cookie{
		Name:     "RefreshToken",
		Value:    refreshTokenString,
		HttpOnly: true,
	}
	http.SetCookie(w, &refreshCookie)
}

func grabCsrfFromRequest(r *http.Request) string {
	csrfFrom := r.FormValue("X-CSRF-Token")

	if csrfFrom != "" {
		return csrfFrom
	} else {

		return r.Header.Get("X-CSRF-Token")
	}
}
