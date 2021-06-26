#!/usr/bin/make

GO111MODULE = on
APP_NAME = go-starter
CMD_PATH = ./cmd/${APP_NAME}
GIT_TAG_HEAD := `git describe --tags --match "v*" --always`
BUILD_TIME := `date +%FT%T%z`

LDFLAGS:=-ldflags "-X main.Version=$(GIT_TAG_HEAD) -X main.BuildTime=$(BUILD_TIME)"

ifeq ($(debug), 1)
LDFLAGS+= -gcflags "-N -l"
endif

.PHONY: ent-init generate build run
ent-init:
ifeq ($(s),)
	@echo "schema is null. e.g: make ent-init s=User"
else
	@ent init --target internal/schema/db $(s)
endif

generate:
	@go generate -x ./...

build-linux:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o ${APP_NAME} ${CMD_PATH}

build-windows:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o ${APP_NAME} ${CMD_PATH}

swag:
	@swag init --parseDependency --parseInternal --parseDepth 1 -g ${CMD_PATH}/main.go -o ./swagger

build:
	@go build $(LDFLAGS) -o ${APP_NAME} ${CMD_PATH}

run:generate
	@go run ${CMD_PATH}
