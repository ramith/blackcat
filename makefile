# File: Makefile

GO             := go
MIN_GO_VERSION := 1.24

MAIN          := main.go
TARGET        := blackcat

PLUGINS_DIR   := plugins
PLUGINS_GO    := $(wildcard $(PLUGINS_DIR)/*.go)
PLUGINS_SO    := $(PLUGINS_GO:.go=.so)

.PHONY: all check_go_version blackcat plugins clean run

all: check_go_version blackcat plugins

check_go_version:
	@echo "Checking Go version >= $(MIN_GO_VERSION)..."
	@version=$$(go version 2>/dev/null | cut -d ' ' -f3 | sed 's/go//' | sed -E 's/([0-9]+\.[0-9]+).*/\1/'); \
	if [ -z "$$version" ]; then \
	  echo "Could not determine Go version. Is Go installed?" && exit 1; \
	fi; \
	lesser=$$(awk 'BEGIN { if ('$$version' < '$(MIN_GO_VERSION)') print 1; else print 0; }'); \
	if [ "$$lesser" -eq 1 ]; then \
	  echo "Go version $$version is older than $(MIN_GO_VERSION). Please upgrade." && exit 1; \
	fi; \
	echo "Go version $$version is OK."

blackcat: check_go_version
	$(GO) build -o $(TARGET) $(MAIN)

plugins: check_go_version $(PLUGINS_SO)

%.so: %.go
	@echo "Building plugin $< → $@"
	$(GO) build -buildmode=plugin -o $@ $<

clean:
	rm -f $(TARGET)
	rm -f $(PLUGINS_DIR)/*.so

run: all
	@echo "Starting $(TARGET)..."
	./$(TARGET)
