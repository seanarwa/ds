FROM ubuntu:18.04

COPY build /app
COPY docker-entrypoint.sh /app/docker-entrypoint.sh

ENTRYPOINT [ "bash /app/docker-entrypoint.sh" ]