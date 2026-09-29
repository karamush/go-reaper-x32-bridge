# Universal entry point: Linux, macOS, and Git Bash / WSL / Cygwin on Windows.
# Recipes are POSIX shell: on Windows run them from Git Bash, or use the
# PowerShell script tools\build.ps1 (which builds the same five binaries).
#
#   make build      build the five commands into bin/
#   make test       unit tests of both modules
#   make check      gofmt + go vet + go test
#   make snapshot   local release dry run into dist/ (needs goreleaser)
#   make gen        regenerate the command catalogue (needs the upstream C sources)
#   make clean      remove bin/ and dist/
#

GO      ?= go
BIN     ?= bin
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS := -s -w -X main.version=$(VERSION)

.PHONY: all build test vet fmt check snapshot gen clean help

all: build

build: | $(BIN)
	$(GO) build -trimpath -ldflags "$(LDFLAGS)" -o $(BIN)/x32emu    ./emulator/cmd/x32emu
	$(GO) build -trimpath -ldflags "$(LDFLAGS)" -o $(BIN)/x32probe  ./emulator/cmd/x32probe
	$(GO) build -trimpath -ldflags "$(LDFLAGS)" -o $(BIN)/x32reaper ./bridge/cmd/x32reaper
	$(GO) build -trimpath -ldflags "$(LDFLAGS)" -o $(BIN)/x32bridge ./bridge/cmd/x32bridge
	$(GO) build -trimpath -ldflags "$(LDFLAGS)" -o $(BIN)/x32listen ./bridge/cmd/x32listen
	@echo "built into $(BIN)/ (version $(VERSION))"

$(BIN):
	mkdir -p $(BIN)

test:
	$(GO) test ./emulator/... ./bridge/...

vet:
	$(GO) vet ./emulator/... ./bridge/...

fmt:
	@unformatted="$$(gofmt -l emulator bridge)"; \
	if [ -n "$$unformatted" ]; then \
		echo "formatting: $$unformatted"; gofmt -w emulator bridge; \
	else \
		echo "gofmt: clean"; \
	fi

check: fmt vet test

# Release dry run: builds every platform into dist/ without publishing anything.
snapshot:
	goreleaser release --clean --snapshot
	@ls -1 dist/*.tar.gz dist/*.zip 2>/dev/null || true

# Regenerates the command catalogue and the Go tables. Needs PowerShell (pwsh on
# Linux/macOS) and a checkout of the upstream C sources: tools/gen_tables.ps1 -Src
gen:
	pwsh -File tools/gen_tables.ps1
	pwsh -File tools/gen_appendix.ps1

clean:
	rm -rf $(BIN) dist

help:
	@sed -n '1,10p' Makefile | sed 's/^# \{0,1\}//'
