FROM golang:alpine

WORKDIR /controle-financeiro-go

ADD . .

RUN go mod download


RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -command="./controle-financeiro-go"