# Find all application directories
APP_DIRS := $(wildcard cmd/*) $(wildcard services/*) $(wildcard playground/*)
# Get just the application names
APPS := $(notdir $(APP_DIRS))

# Default application to run if not specified
APP ?= demo-http-server

# Find the directory for the selected application. Use = for deferred expansion.
APP_DIR = $(shell find cmd services playground -maxdepth 1 -type d -name "$(APP)" | head -n 1)

# Go parameters
GOBASE := $(shell pwd)
GOPATH := $(GOBASE)/.bin
GOBIN  := $(GOPATH)/bin
GOFILES := $(wildcard *.go)

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage: make <target> [APP=<app_name>]'
	@echo 'Targets:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
	@echo 'Applications:'
	@for app in $(APPS); do \
		echo "  $$app"; \
	done


.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	@test -z "$(shell git status --porcelain)"


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: run all quality control checks
.PHONY: audit
audit:
	@echo "--- Running Quality Control Checks ---"
	go mod tidy -diff
	go mod verify
	test -z "$(shell gofmt -l . | grep -v 'tools/tools.go')"
	go vet ./...
	go build ./...
	go run honnef.co/go/tools/cmd/staticcheck -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck ./...
	@echo "--- Quality Control Checks Passed ---"

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## upgradeable: list direct dependencies that have upgrades available
.PHONY: upgradeable
upgradeable:
	@go run github.com/oligot/go-mod-upgrade @latest

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## tidy: tidy modfiles and format .go files
.PHONY: tidy
tidy:
	go mod tidy -v
	go fmt ./...

## build: build the application
.PHONY: build
build:
	@if [ -z "$(APP_DIR)" ]; then \
		echo "Error: Application '$(APP)' not found." >&2; \
		exit 1; \
	fi
	@echo "Building $(APP) from ./$(APP_DIR)..."
	@go build -o=/tmp/bin/$(APP) ./$(APP_DIR)

## run: run the application
.PHONY: run
run: build
	@/tmp/bin/$(APP)

## run/live: run the application with reloading on file changes
.PHONY: run/live
run/live:
	@if [ -z "$(APP_DIR)" ]; then \
		echo "Error: Application '$(APP)' not found." >&2; \
		exit 1; \
	fi
	@go run github.com/air-verse/air --build.cmd "make build APP=$(APP)" --build.bin "/tmp/bin/$(APP)"


# ==================================================================================== #
# OPERATIONS
# ==================================================================================== #

## push: push changes to the remote Git repository
.PHONY: push
push: confirm audit no-dirty
	git push

## production/deploy: deploy the application to production
.PHONY: production/deploy
production/deploy: confirm audit no-dirty
	@if [ -z "$(APP_DIR)" ]; then \
		echo "Error: Application '$(APP)' not found." >&2; \
		exit 1; \
	fi
	@echo "Building $(APP) for production..."
	@GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=/tmp/bin/linux_amd64/$(APP) ./$(APP_DIR)
	# upx -5 /tmp/bin/linux_amd64/$(APP)
	# Include additional deployment steps here...


# ==================================================================================== #
# OTHER
# ==================================================================================== #

## change-username: run the username changer tool
.PHONY: change-username
change-username:
	@go run ./tools/username-changer/main.go

## clean: remove temporary files
.PHONY: clean
clean:
	rm -rf /tmp/bin
	rm -f /tmp/coverage.out

## self-update: update this Makefile from a template
.PHONY: self-update
	@echo "Not implemented"

## version: print the version of the application
.PHONY: version
version:
	@echo "Not implemented"