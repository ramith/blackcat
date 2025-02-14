package main

import (
    "fmt"
    "net/http"
    "path/filepath"
    "plugin"
    "strings"
    "sync"
    "time"

    "github.com/ramith/blackcat/pkg/iface"
)

var (
    pluginLock sync.Mutex
    plugins    = map[string]iface.Plugin{}
)

func loadPlugin(path string) error {
    mod, err := plugin.Open(path)
    if err != nil {
        return err
    }
    sym, err := mod.Lookup("PluginInstance")
    if err != nil {
        return err
    }
    inst, ok := sym.(iface.Plugin)
    if !ok {
        return fmt.Errorf("invalid plugin symbol in %s", path)
    }

    pluginLock.Lock()
    defer pluginLock.Unlock()
    plugins[inst.Name()] = inst
    fmt.Printf("Loaded plugin: %s\n", inst.Name())
    return nil
}

func watchPlugins(dir string) {
    known := map[string]struct{}{}
    for {
        soFiles, _ := filepath.Glob(filepath.Join(dir, "*.so"))
        for _, file := range soFiles {
            if _, exists := known[file]; !exists {
                if err := loadPlugin(file); err == nil {
                    known[file] = struct{}{}
                } else {
                    fmt.Printf("Failed to load %s: %v\n", file, err)
                }
            }
        }
        time.Sleep(5 * time.Second)
    }
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    pluginLock.Lock()
    defer pluginLock.Unlock()

    for _, p := range plugins {
        if strings.HasPrefix(r.URL.Path, p.Prefix()) {
            http.StripPrefix(p.Prefix(), p.Handler()).ServeHTTP(w, r)
            return
        }
    }
    http.NotFound(w, r)
}

func main() {
    go watchPlugins("./plugins")
    http.HandleFunc("/", handleRequest)
    fmt.Println("Blackcat server started on :8080")
    http.ListenAndServe(":8080", nil)
}
