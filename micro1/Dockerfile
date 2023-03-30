# GRPC-demo
FROM golang:latest

# Install protoc and zip system library
RUN apt-get update && apt-get install -y zip && \
    mkdir /opt/protoc && cd /opt/protoc && wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.0/protoc-3.7.0-linux-x86_64.zip && \
    unzip protoc-3.7.0-linux-x86_64.zip

ENV PATH=$PATH:$GOPATH/bin:/opt/protoc/bin

COPY config/config.go /app/micro1/config/config.go
COPY config/config.yml /app/config.yml
COPY internal /app/micro1
COPY pkg /app/micro1
COPY proto /app/micro1
COPY server /app/micro1

WORKDIR /app

RUN go mod init github.com/bav-demo
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.54
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3
RUN go mod tidy

RUN cd /app && \
    protoc proto/*.proto --go_out=:. --go-grpc_out=:.

ENTRYPOINT cd /app/server && go run main.go
