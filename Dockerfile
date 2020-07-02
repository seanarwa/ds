# =============================================================================
# build stage
#
# install golang dependencies & build binaries
# =============================================================================
FROM golang:1.14 AS build

ENV DS_BUILD_DIR=/bin/build
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=off

WORKDIR /go/src/github.com/seanarwa/ds

COPY . .
RUN mkdir -p $DS_BUILD_DIR
RUN go get -v -t -d ./...
RUN go build -o $DS_BUILD_DIR -v .

# =============================================================================
# final stage
#
# add static assets and copy binaries from build stage
# =============================================================================
FROM alpine:3.12

RUN addgroup -S ds && adduser -S ds ds
USER ds

WORKDIR /app

COPY --from=build $DS_BUILD_DIR .
COPY ./conf/ ./conf/
COPY docker-entrypoint.sh ./docker-entrypoint.sh

RUN ls

ENTRYPOINT [ "/bin/bash", "-c", "./docker-entrypoint.sh" ]