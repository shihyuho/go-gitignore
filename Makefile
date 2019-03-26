HAS_GOLINT := $(shell command -v golint;)
VERSION :=
DIST := $(CURDIR)/_dist
BUILD := $(CURDIR)/_build
BINARY := gitignore
MAIN := ./cmd/gitignore
LICENSE := ./LICENSE

.PHONY: test
test: golint
	go test ./... -v

.PHONY: gofmt
gofmt:
	gofmt -s -w .

.PHONY: golint
golint: gofmt
ifndef HAS_GOLINT
	go get -u golang.org/x/lint/golint
endif
	golint -set_exit_status ./cmd/...
	golint -set_exit_status ./pkg/...

.PHONY: build
build: clean bootstrap
	mkdir -p $(BUILD)
	go build -o $(BUILD)/$(BINARY) $(MAIN)

.PHONY: dist
dist:
ifeq ($(strip $(VERSION)),)
	$(error VERSION is not set)
endif
	go get -u github.com/inconshreveable/mousetrap
	mkdir -p $(BUILD)
	mkdir -p $(DIST)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD)/$(BINARY) -ldflags $(LDFLAGS) -a -tags netgo $(MAIN)
	tar -C $(BUILD) -zcvf $(DIST)/$(BINARY)-linux-$(VERSION).tgz $(BINARY) $(LICENSE)
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD)/$(BINARY) -ldflags $(LDFLAGS) -a -tags netgo $(MAIN)
	tar -C $(BUILD) -zcvf $(DIST)/$(BINARY)-darwin-$(VERSION).tgz $(BINARY) $(LICENSE)
	GOOS=windows GOARCH=amd64 go build -o $(BUILD)/$(BINARY).exe -ldflags $(LDFLAGS) -a -tags netgo $(MAIN)
	tar -C $(BUILD) -llzcvf $(DIST)/$(BINARY)-windows-$(VERSION).tgz $(BINARY).exe $(LICENSE)

.PHONY: bootstrap
bootstrap:
ifeq (,$(wildcard ./go.mod))
	go mod init gitignore
endif
	go mod download

.PHONY: clean
clean:
	rm -rf _*
