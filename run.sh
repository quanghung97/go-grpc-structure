# golang server generate
protoc pb/echo.proto --go_out=:. --go-grpc_out=:.
# js client generate
protoc pb/echo.proto --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:.
# go run server/server.go
# py -m http.server 8081 &
#go run client/client.go
