FROM golang:latest

RUN mkdir /app
WORKDIR /app
ADD . /app


RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

ENV PORT=8080

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main



