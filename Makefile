# Copyright (C) 2015, 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

APP = enigma
EXE = bin/enigma

SHELL = /bin/bash

DIR = $(shell pwd)

DOCKER = docker

GB = gb

GOX = gox -os="linux darwin windows freebsd openbsd netbsd"

BINTRAY_URI = https://api.bintray.com
BINTRAY_USERNAME = nlamirault
BINTRAY_REPOSITORY= oss

NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

SRC = src/github.com/nlamirault/enigma

SRCS = $(shell git ls-files '*.go' | grep -v '^vendor/')
PKGS = $(shell find src -type f -print0 | xargs -0 -n 1 dirname | sort -u|sed -e "s/^src\///g")
EXE = $(shell ls enigma_*)

VERSION=$(shell \
        grep "const Version" $(SRC)/version/version.go \
        |awk -F'=' '{print $$2}' \
        |sed -e "s/[^0-9.]//g" \
	|sed -e "s/ //g")

PACKAGE=$(APP)-$(VERSION)
ARCHIVE=$(PACKAGE).tar

all: help

help:
	@echo -e "$(OK_COLOR)==== $(APP) [$(VERSION)] ====$(NO_COLOR)"
	@echo -e "$(WARN_COLOR)init$(NO_COLOR)      :  Install requirements"
	@echo -e "$(WARN_COLOR)build$(NO_COLOR)     :  Make all binaries"
	@echo -e "$(WARN_COLOR)test$(NO_COLOR)      :  Launch unit tests"
	@echo -e "$(WARN_COLOR)lint$(NO_COLOR)      :  Launch golint"
	@echo -e "$(WARN_COLOR)vet$(NO_COLOR)       :  Launch go vet"
	@echo -e "$(WARN_COLOR)coverage$(NO_COLOR)  :  Launch code coverage"
	@echo -e "$(WARN_COLOR)clean$(NO_COLOR)     :  Cleanup"
	@echo -e "$(WARN_COLOR)binaries$(NO_COLOR)  :  Make binaries"

clean:
	@echo -e "$(OK_COLOR)[$(APP)] Cleanup$(NO_COLOR)"
	@rm -fr $(EXE) $(APP)-*.tar.gz pkg bin $(APP)_*

.PHONY: init
init:
	@echo -e "$(OK_COLOR)[$(APP)] Install requirements$(NO_COLOR)"
	@go get -u github.com/golang/glog
	@go get -u github.com/constabulary/gb/...
	@go get -u github.com/golang/lint/golint
	@go get -u github.com/kisielk/errcheck
	@go get -u golang.org/x/tools/cmd/oracle
	@go get -u github.com/mitchellh/gox

.PHONY: build
build:
	@echo -e "$(OK_COLOR)[$(APP)] Build $(NO_COLOR)"
	@$(GB) build all

.PHONY: test
test:
	@echo -e "$(OK_COLOR)[$(APP)] Launch unit tests $(NO_COLOR)"
	@$(GB) test all -test.v=true

.PHONY: lint
lint:
	@$(foreach file,$(SRCS),golint $(file) || exit;)

.PHONY: vet
vet:
	@$(foreach file,$(SRCS),go vet $(file) || exit;)

.PHONY: errcheck
errcheck:
	@echo -e "$(OK_COLOR)[$(APP)] Go Errcheck $(NO_COLOR)"
	@$(foreach pkg,$(PKGS),env GOPATH=`pwd`:`pwd`/vendor errcheck $(pkg) || exit;)

.PHONY: coverage
coverage:
	@$(foreach pkg,$(PKGS),env GOPATH=`pwd`:`pwd`/vendor go test -cover $(pkg) || exit;)

gox:
	@echo -e "$(OK_COLOR)[$(APP)] Create binaries $(NO_COLOR)"
	$(GOX) github.com/nlamirault/enigma

.PHONY: binaries
binaries: gox
	@echo -e "$(OK_COLOR)[$(APP)] Upload binaries to Bintray $(NO_COLOR)"
	for i in $(EXE); do \
		curl -T $$i \
			-u$(BINTRAY_USERNAME):$(BINTRAY_APIKEY) \
			"$(BINTRAY_URI)/content/$(BINTRAY_USERNAME)/$(BINTRAY_REPOSITORY)/$(APP)/${VERSION}/$$i;publish=1"; \
        done


# for goprojectile
.PHONY: gopath
gopath:
	@echo GOPATH=`pwd`:`pwd`/vendor
