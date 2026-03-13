package api

import (
	"net/http"
)

func (a *API) setupRoutes() {
	a.router.Get("/", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
}
