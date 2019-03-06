# Go tools.
GO ?= GO111MODULE=off go
GO_MD2MAN ?= go-md2man

# Paths.
PROJECT := github.com/SUSE/gitguy-finglonger
CMD := .

# We use Docker because Go is just horrific to deal with.
GITHUBHOOK_IMAGE := gitguy-finglonger_dev
DOCKER_RUN := docker run --rm -it --security-opt apparmor:unconfined --security-opt label:disable -v ${PWD}:/go/src/${PROJECT}

# Output directory.
BUILD_DIR ?= ./bin

# Release information.
GPG_KEYID ?=

# Version information.
VERSION := $(shell cat VERSION)
SECURITY_TOKEN := $(shell echo ${GITHUB_TOKEN_WEBHOOK})
SECRET_TOKEN := $(shell echo ${GITHUB_SECRET_WEBHOOK})
COMMIT_NO := $(shell git rev-parse HEAD 2> /dev/null || true)
COMMIT := $(if $(shell git status --porcelain --untracked-files=no),"${COMMIT_NO}-dirty","${COMMIT_NO}")

# Get current Version changelog
CHANGE := $(shell sed -e '1,/$(VERSION)/d;/v.*/Q' ./CHANGELOG.md)

BUILD_FLAGS ?=

BASE_FLAGS := ${BUILD_FLAGS} -tags "${BUILDTAGS}"

BASE_LDFLAGS := -X $(PROJECT)/pkg/github/api.version=$(VERSION)

DYN_BUILD_FLAGS := ${BASE_FLAGS} -buildmode=pie -ldflags "${BASE_LDFLAGS}"
TEST_BUILD_FLAGS := ${BASE_FLAGS} -buildmode=pie -ldflags "${BASE_LDFLAGS} -X ${PROJECT}/pkg/testutils.binaryType=test"
STATIC_BUILD_FLAGS := ${BASE_FLAGS} -ldflags "${BASE_LDFLAGS} -extldflags '-static'"

.DEFAULT: gitguy-finglonger

GO_SRC = $(shell find . -name \*.go)

# NOTE: If you change these make sure you also update local-validate-build.

gitguy-finglonger: $(GO_SRC)
	$(GO) build ${DYN_BUILD_FLAGS} -o $(BUILD_DIR)/$@ ${CMD}

gitguy-finglonger.static: $(GO_SRC)
	env CGO_ENABLED=0 $(GO) build ${STATIC_BUILD_FLAGS} -o $(BUILD_DIR)/$@ ${CMD}

install: $(GO_SRC)
	$(GO) install -v ${DYN_BUILD_FLAGS} ${CMD}

install.static: $(GO_SRC)
	$(GO) install -v ${STATIC_BUILD_FLAGS} ${CMD}

clean:
	rm -rf ./bin ./build ./release ./.cover .tmp-validate coverage.txt
	rm -f $(MANPAGES)

local-validate: local-validate-git local-validate-go local-validate-reproducible

EPOCH_COMMIT ?= d51616f774e691623dc40df0ae5f90c35b0e1fd0
local-validate-git:
	@type git-validation > /dev/null 2>/dev/null || (echo "ERROR: git-validation not found." && false)
ifdef TRAVIS_COMMIT_RANGE
	git-validation -q -run DCO,short-subject
else
	git-validation -q -run DCO,short-subject -range $(EPOCH_COMMIT)..HEAD
endif

local-validate-go:
	@type gofmt    >/dev/null 2>/dev/null || (echo "ERROR: gofmt not found." && false)
	test -z "$$(gofmt -s -l . | grep -vE '^vendor/|^third_party/' | tee /dev/stderr)"
	@type golint   >/dev/null 2>/dev/null || (echo "ERROR: golint not found." && false)
	test -z "$$(golint ./... | grep -v -E '^vendor/|^third_party/' | tee /dev/stderr)"
	@go doc cmd/vet >/dev/null 2>/dev/null || (echo "ERROR: go vet not found." && false)
	test -z "$$($(GO) vet $$($(GO) list $(PROJECT)/... | grep -vE '/vendor/|/third_party/') 2>&1 | tee /dev/stderr)"

# Make sure that our builds are reproducible even if you wait between them and
# the modified time of the files is different.
local-validate-reproducible:
	mkdir -p .tmp-validate
	make -B gitguy-finglonger && cp $(BUILD_DIR)/gitguy-finglonger .tmp-validate/gitguy-finglonger.a
	@echo sleep 10s
	@sleep 10s && touch $(GO_SRC)
	make -B gitguy-finglonger && cp $(BUILD_DIR)/gitguy-finglonger .tmp-validate/gitguy-finglonger.b
	diff -s .tmp-validate/gitguy-finglonger.{a,b}
	sha256sum .tmp-validate/gitguy-finglonger.{a,b}
	rm -r .tmp-validate/gitguy-finglonger.{a,b}

local-validate-build:
	$(GO) build ${DYN_BUILD_FLAGS} -o /dev/null ${CMD}
	env CGO_ENABLED=0 $(GO) build ${STATIC_BUILD_FLAGS} -o /dev/null ${CMD}
	$(GO) test -run nothing ${DYN_BUILD_FLAGS} $(PROJECT)/...

# Used for tests.
DOCKER_IMAGE :=kubic-project/amd64:tumbleweed

gitguy-finglonger-image:
	docker build -t $(GITHUBHOOK_IMAGE) .


test.unit: gitguy-finglonger-image
	$(DOCKER_RUN) $(GITHUBHOOK_IMAGE) make test

test: local-validate-go
	rm -rf /tmp/gitguy-finglonger
	$(GO) test -v ./...

cover:
	bash <scripts/cover.sh

dist: export COPYFILE_DISABLE=1 #teach OSX tar to not put ._* files in tar archive
dist:
	rm -rf build/gitguy-finglonger/* release/*
	mkdir -p build/gitguy-finglonger/bin release/
	cp README.md LICENSE build/gitguy-finglonger
	GOOS=linux GOARCH=amd64 go build -o build/gitguy-finglonger/bin/gitguy-finglonger -ldflags="$(BASE_LDFLAGS)"
	tar -C build/ -zcvf $(CURDIR)/release/gitguy-finglonger-linux.tgz gitguy-finglonger/
	GOOS=darwin GOARCH=amd64 go build -o build/gitguy-finglonger/bin/gitguy-finglonger -ldflags="$(BASE_LDFLAGS)"
	tar -C build/ -zcvf $(CURDIR)/release/gitguy-finglonger-macos.tgz gitguy-finglonger/
	rm build/gitguy-finglonger/bin/gitguy-finglonger

release: dist
ifndef GITHUB_TOKEN
	$(error GITHUB_TOKEN is undefined)
endif
	github-release release -u kubic-project -r gitguy-finglonger --tag $(VERSION)  --name $(VERSION) -s $(GITHUB_TOKEN) -d "$(CHANGE)"
	github-release upload -u kubic-project -r gitguy-finglonger --tag $(VERSION)  --name gitguy-finglonger-linux.tgz --file release/gitguy-finglonger-linux.tgz -s $(GITHUB_TOKEN)
	github-release upload -u kubic-project -r gitguy-finglonger --tag $(VERSION)  --name gitguy-finglonger-macos.tgz --file release/gitguy-finglonger-macos.tgz -s $(GITHUB_TOKEN)
	github-release upload -u kubic-project -r gitguy-finglonger --tag $(VERSION)  --name gitguy-finglonger-windows.tgz --file release/gitguy-finglonger-windows.tgz -s $(GITHUB_TOKEN)

MANPAGES_MD := $(wildcard doc/man/*.md)
MANPAGES    := $(MANPAGES_MD:%.md=%)

doc/man/%.1: doc/man/%.1.md
	@$(GO_MD2MAN) -in $< -out $@.out
	@$(GO) run doc/man/sanitize.go $@.out &> $@
	@rm $@.out

doc: $(MANPAGES)

.PHONY: gitguy-finglonger \
	gitguy-finglonger.static \
	install \
	install.static \
	clean \
	local-validate \
	local-validate-git \
	local-validate-go \
	local-validate-reproducible \
	local-validate-build \
	gitguy-finglonger-image \
	test.unit
	test \
	cover \
	bootstrap \
	dist \
	release \
	doc