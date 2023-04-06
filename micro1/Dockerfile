FROM golang:latest

# Install protoc and zip system library
RUN apt-get update && apt-get install -y zip && \
    mkdir /opt/protoc && cd /opt/protoc && wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.0/protoc-3.7.0-linux-x86_64.zip && \
    unzip protoc-3.7.0-linux-x86_64.zip

ENV PATH=$PATH:$GOPATH/bin:/opt/protoc/bin

COPY config /app/config
COPY internal /app/internal
COPY pkg /app/pkg
COPY proto /app/proto
COPY server /app/server
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum
COPY config.docker.yml /app/config.yml

WORKDIR /app

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3
RUN go mod tidy

RUN protoc proto/authenticate/authentication.proto --go_out=:. --go-grpc_out=:.

RUN protoc proto/society/society.proto --go_out=:. --go-grpc_out=:.

ENTRYPOINT go run server/main.go

# RUN go install github.com/cosmtrek/air@latest
# Command to run the executable
# CMD ["air", "-c", ".air.toml"]
