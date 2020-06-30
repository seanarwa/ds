FROM golang:1.14

USER root

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
# RUN go install -v ./...
RUN go build -v .

ENTRYPOINT [ "/bin/bash", "-c", "./docker-entrypoint.sh" ]