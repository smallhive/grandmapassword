GO_PLATFORM_FLAGS=GOOS=linux GOARCH=amd64
GO_CMD=CGO_ENABLED=0 GO111MODULE=on go
GOPATH?=$(HOME)/go

BUILD_PATH=./build/bin
APP=grandmapassword

LDFLAGS:=-s -w

clean:
	@rm -f ${BUILD_PATH}/${APP}

build: clean
	$(GO_PLATFORM_FLAGS) $(GO_CMD) build -ldflags "$(LDFLAGS)" -o ${BUILD_PATH}/${APP} ./cmd/${APP}

run:
	${BUILD_PATH}/${APP}

test:
	go test ./...

lint:
	golangci-lint run
