
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go mod init github.com/<organization-name>/<project-name>
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: proto
proto:
	protoc micro1/proto/authenticate/authenticate.proto --go_out=:. --go-grpc_out=:.
	