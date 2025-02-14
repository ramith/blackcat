package iface

import "net/http"

// Plugin is the shared interface implemented by every plugin.
type Plugin interface {
    Name() string
    Prefix() string
    Handler() http.Handler
}
