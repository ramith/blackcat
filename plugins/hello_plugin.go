// File: plugins/hello_plugin.go
package main

import (
    "net/http"

    "github.com/ramith/blackcat/pkg/iface" // Adjust import to your module path
)

type helloPlugin struct{}

func (h *helloPlugin) Name() string {
    return "HelloPlugin"
}

// Prefix returns "/hello" so any request under /hello triggers this plugin.
func (h *helloPlugin) Prefix() string {
    return "/hello"
}

// Handler returns a simple handler that responds with "Hello!".
func (h *helloPlugin) Handler() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello!\n"))
    })
}

// PluginInstance is the exported symbol that Blackcat loads via plugin.Lookup.
var PluginInstance iface.Plugin = &helloPlugin{}
