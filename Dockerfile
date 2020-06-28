FROM ubuntu:18.04

USER root

WORKDIR /app

COPY build/ .
COPY docker-entrypoint.sh ./docker-entrypoint.sh

ENTRYPOINT [ "/bin/bash", "-c", "./docker-entrypoint.sh" ]