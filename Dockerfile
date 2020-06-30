FROM alpine:3.12

USER root

WORKDIR /app
COPY build /app

COPY build/ .
COPY docker-entrypoint.sh ./docker-entrypoint.sh

ENTRYPOINT [ "/bin/bash", "-c", "./docker-entrypoint.sh" ]