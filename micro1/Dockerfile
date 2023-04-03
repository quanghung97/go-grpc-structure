# GRPC-demo
FROM golang:latest

# Install protoc and zip system library
RUN apt-get update && apt-get install -y zip && \
    mkdir /opt/protoc && cd /opt/protoc && wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.0/protoc-3.7.0-linux-x86_64.zip && \
    unzip protoc-3.7.0-linux-x86_64.zip

ENV PATH=$PATH:$GOPATH/bin:/opt/protoc/bin

COPY config /app/micro1/config
# COPY config/config_docker.yml /app/config.yml
COPY internal /app/micro1/internal
COPY pkg /app/micro1/pkg
COPY proto /app/micro1/proto
COPY server /app/micro1/server
COPY .env /app/.env

WORKDIR /app

RUN go mod init github.com/bav-demo
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go mod tidy

RUN cd /app && \
    protoc micro1/proto/authenticate/authentication.proto --go_out=:. --go-grpc_out=:.

RUN cd /app && \
    protoc micro1/proto/society/society.proto --go_out=:. --go-grpc_out=:.

ENTRYPOINT cd /app && go run micro1/server/main.go
