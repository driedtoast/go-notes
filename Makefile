UNAME=$(shell uname -s)
SOURCE_DIR=$(shell pwd)/src
PROJECT_NAME=$(shell basename `pwd`)
DEPS_FILE=$(SOURCE_DIR)/../dependencies
DEPENDENCIES=$(shell cat ${DEPS_FILE})
SOURCE_PKGS=$(shell ls ${SOURCE_DIR})

ifndef VERBOSE
MAKEFLAGS+=--no-print-directory
endif

ifeq ($(UNAME),Darwin)
ECHO=echo
else
ECHO=echo -e
endif

ASCIIDOC=asciidoc
CAT=cat

all: build
	-@$(ECHO) "\n\033[1;32mCONGRATULATIONS!\033[0;32m\n$(PROJECT_NAME) has been built and tested!\033[0m\n"

check: build format
	-@$(ECHO) ""

clean:
	@rm -f bin/*
	@rm -Rf pkg/*
	@$(foreach dep,$(DEPENDENCIES), rm -Rf src/$(dep);)

install: build install-packages install-man
	-@$(ECHO) "\n\033[1;32mCONGRATULATIONS!\033[0;32m\n$(PROJECT_NAME) has been built, tested and installed!\033[0m\n"

format:
	-@$(ECHO) "\n\033[0;35m%%% Formatting\033[0m"
	@go fmt ./...

resolve-deps: 
	-@$(ECHO) "\033[0;35m%%% Resolving dependencies for $(PROJECT_NAME)\033[0m"
	@$(foreach dep,$(DEPENDENCIES), go get -v $(dep);)

build: resolve-deps clean $(BUILD_MAN)
	-@$(ECHO) "Using GOPATH $(GOPATH)"
	-@$(ECHO) "\n\033[0;35m%%% Building $(PROJECT_NAME)\033[0m"
	# TODO debug exclusions
	@$(foreach pkg,$(SOURCE_PKGS), $(if !$(findstring $(pkg),$(DEPENDENCIES)),go build -v -o ./bin/$(pkg) $(pkg);))

test: build
	-@$(ECHO) "\n\033[0;35m%%% Running tests\033[0m"
	@go test ./...

install-packages:
	-@$(ECHO) "\n\033[0;36m%%% Installing project\033[0m"
	@$(foreach pkg,$(SOURCE_PKGS), go install -v $(pkg);)

man: clean-man
	@$(MAKE) -C docs

install-man:
	-@$(ECHO) "\033[0;36mInstalling documentation\033[0m"
	@$(MAKE) -C docs install

clean-man:
	-@$(MAKE) -C docs clean