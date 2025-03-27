.PHONY: all clean build-plugin run-server

API_PATH=./api
PLUGIN_PATH=./plugins/sampleplugin
SERVER_PATH=./server

PLUGIN_SO=$(PLUGIN_PATH)/plugin.so

all: build-plugin run-server

build-plugin:
	cd $(PLUGIN_PATH) && \
	go mod tidy && \
	go build -buildmode=plugin -o plugin.so restapi.go

run-server:
	cd $(SERVER_PATH) && \
	go mod tidy && \
	go run main.go

clean:
	rm -f $(PLUGIN_SO)
	go clean -cache -modcache
