package main

import (
	"net/http"
	"github.com/ramith/blackcat/api"
)

type PluginImpl struct{}

func (p *PluginImpl) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from plugin"))
	})
}

var Handler api.Plugin = &PluginImpl{}
