package api

import "net/http"

type Plugin interface {
	RegisterRoutes(mux *http.ServeMux)
}
