
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go mod init github.com/<organization-name>/<project-name>
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: proto
proto:
	protoc micro1/proto/authenticate/authenticate.proto --go_out=:. --go-grpc_out=:.
	
# .PHONY: build
# build:
# 	go build -o example *.go

# .PHONY: test
# test:
# 	go test -v ./... -cover

# .PHONY: docker
# docker:
# 	docker build . -t example:latest


# .PHONY:

# local:
# 	echo "Starting local docker compose"
# 	docker-compose -f docker-compose.local.yml up --build

# gen:
# 	GO111MODULE=on  swagger generate spec -o ./api/swagger/swagger.yaml --scan-models

# upload:
# 	sudo docker build -t quanghung97/products_microservice:latest -f ./Dockerfile .
# 	sudo docker push quanghung97/products_microservice:latest
# 	#sudo APP_VERSION=latest docker-compose up

# pull:
# 	sudo docker pull quanghung97/products_microservice:latest


# # ==============================================================================
# # Modules support

# deps-reset:
# 	git checkout -- go.mod
# 	go mod tidy
# 	go mod vendor

# tidy:
# 	go mod tidy
# 	go mod vendor

# deps-upgrade:
# 	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
# 	go get -u -t -d -v ./...
# 	go mod tidy
# 	go mod vendor

# deps-cleancache:
# 	go clean -modcache

# # ==============================================================================
# # Linters

# run-linter:
# 	echo "Starting linters"
# 	golangci-lint run ./...


# # ==============================================================================
# # Docker support

# FILES := $(shell docker ps -aq)

# down-local:
# 	docker stop $(FILES)
# 	docker rm $(FILES)

# clean:
# 	docker system prune -f

# logs-local:
# 	docker logs -f $(FILES)
