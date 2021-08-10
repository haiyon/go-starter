#!/usr/bin/make
GO111MODULE = on
APP_NAME = go-starter
CMD_PATH = ./cmd/${APP_NAME}
OUT = ./out
GIT_TAG_HEAD := `git describe --tags --match "v*" --always`
BUILD_TIME := `date +%FT%T%z`

LDFLAGS:=-ldflags "-X main.Version=$(GIT_TAG_HEAD) -X main.BuildTime=$(BUILD_TIME)"

ifeq ($(debug), 1)
LDFLAGS+= -gcflags "-N -l"
endif

.PHONY: ent-init generate build-linux build run optimize
ent-init:
ifeq ($(s),)
	@echo "schema is null. e.g: make ent-init s=User"
else
	@ent init --target internal/schema/db $(s)
endif

generate:
	@go generate -x ./...

build-linux:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o ${OUT}/${APP_NAME} ${CMD_PATH}
	if [ ! -d "${OUT}" ]; then mkdir ${OUT}; fi
	if [ ! -f "${OUT}/config.yml" ]; then cp -r ./configs/config.yml ${OUT}; fi

build:
	@go build $(LDFLAGS) -o ${OUT}/${APP_NAME} ${CMD_PATH}

swag:
	@swag init --parseDependency --parseInternal --parseDepth 1 -g ${CMD_PATH}/main.go -o ./swagger

run:
	@go run ${CMD_PATH}

optimize:build-linux
	@upx -9 ${OUT}/${APP_NAME}
