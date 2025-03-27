package main

import (
	"fmt"
	"net/http"
	"os"
	"plugin"

	"gopkg.in/yaml.v3"

	"github.com/ramith/blackcat/api"
)

type Config struct {
	PluginPath    string `yaml:"plugin_path"`
	HandlerSymbol string `yaml:"handler_symbol"`
}

func main() {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(fmt.Errorf("failed to read config: %w", err))
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(fmt.Errorf("invalid config format: %w", err))
	}

	if _, err := os.Stat(cfg.PluginPath); os.IsNotExist(err) {
		panic(fmt.Errorf("plugin not found: %s", cfg.PluginPath))
	}

	p, err := plugin.Open(cfg.PluginPath)
	if err != nil {
		panic(fmt.Errorf("failed to open plugin: %w", err))
	}

	sym, err := p.Lookup(cfg.HandlerSymbol)
	if err != nil {
		panic(fmt.Errorf("symbol not found: %w", err))
	}

	pluginHandler, ok := sym.(api.Plugin)
	if !ok {
		panic("plugin does not implement api.Plugin interface")
	}

	mux := http.NewServeMux()
	pluginHandler.RegisterRoutes(mux)

	fmt.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(fmt.Errorf("server error: %w", err))
	}
}
