# syntax=docker/dockerfile:1

FROM golang:1.16

WORKDIR /tbc

RUN apt-get update
RUN apt-get install -y protobuf-compiler
RUN go get -u -v github.com/golang/protobuf/proto
RUN go get -u -v github.com/golang/protobuf/protoc-gen-go

RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh | bash

ENV NODE_VERSION=14.17.6
ENV NVM_DIR="/root/.nvm"
RUN . "$NVM_DIR/nvm.sh" && nvm install ${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm use v${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm alias default v${NODE_VERSION}
ENV PATH="/root/.nvm/versions/node/v${NODE_VERSION}/bin/:${PATH}"

EXPOSE 8080/tcp
