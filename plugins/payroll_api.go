package main

import (
    "net/http"
    "github.com/ramith/blackcat/pkg/iface"
)

type payrollPlugin struct {
    mux *http.ServeMux
}

func (p *payrollPlugin) Name() string {
    return "Payroll API"
}

func (p *payrollPlugin) Prefix() string {
    return "/payroll/"
}

func (p *payrollPlugin) Handler() http.Handler {
    return p.mux
}

// Exported symbol must be named PluginInstance
var PluginInstance iface.Plugin = newPayrollPlugin()

func newPayrollPlugin() iface.Plugin {
    mux := http.NewServeMux()

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Payroll root endpoint\n"))
    })

    mux.HandleFunc("/x", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Payroll subpath X\n"))
    })

    return &payrollPlugin{mux: mux}
}
