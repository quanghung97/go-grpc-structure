
FROM node:10

# Install protoc and zip system library
RUN apt-get update && apt-get install -y zip && \
    mkdir /opt/protoc && cd /opt/protoc && wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.0/protoc-3.7.0-linux-x86_64.zip && \
    unzip protoc-3.7.0-linux-x86_64.zip && \
    cd /opt/protoc/bin && wget https://github.com/grpc/grpc-web/releases/download/1.4.2/protoc-gen-grpc-web-1.4.2-linux-x86_64 && \
    cd /opt/protoc/bin && mv protoc-gen-grpc-web-1.4.2-linux-x86_64 protoc-gen-grpc-web && \
    chmod +x /opt/protoc/bin/protoc-gen-grpc-web
# install python3 to host static file
RUN apt-get install python -y

ENV PATH=$PATH:$GOPATH/bin:/opt/protoc/bin


# Copy the project to be executed
RUN mkdir -p /js/src/grpc-demo 
COPY client/client.js /js/src/grpc-demo/client/client.js
COPY index.html /js/src/grpc-demo/index.html
COPY package.json /js/src/grpc-demo/package.json
COPY pb/echo.proto /js/src/grpc-demo/pb/echo.proto

WORKDIR /js/src/grpc-demo
RUN npm install

# Gen js-out and grpc-web_out
RUN cd /js/src/grpc-demo/pb && \
    protoc echo.proto --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. 

RUN npm install -g webpack webpack-cli
RUN npx webpack ./client/client.js

ENTRYPOINT cd /js/src/grpc-demo && python3 -m http.server 8182 
