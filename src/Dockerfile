# GRPC-demo
# build command : docker build src/. -t grpc-docker/server
# run command : docker run -it grpc-docker/server
FROM golang:latest

# Install grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# Install protoc and zip system library
RUN apt-get update && apt-get install -y zip && \
    mkdir /opt/protoc && cd /opt/protoc && wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.0/protoc-3.7.0-linux-x86_64.zip && \
    unzip protoc-3.7.0-linux-x86_64.zip

ENV PATH=$PATH:$GOPATH/bin:/opt/protoc/bin

COPY go.mod /go/src
COPY go.sum /go/src

# Copy the grpc proto file and generate the go module
RUN mkdir -p /go/src/pb
COPY pb/echo.proto /go/src/pb

RUN cd /go/src && \
    protoc pb/echo.proto --go_out=:. --go-grpc_out=:.

# Copy the project to be executed
RUN mkdir -p /go/src/grpc-server 
COPY server/main.go /go/src/grpc-server

ENTRYPOINT cd /go/src/grpc-server && go run main.go
