FROM alpine:latest


WORKDIR /app/server
COPY ./server /app/server/
COPY ./config.dev.yaml /app/server/

EXPOSE 8888
ENTRYPOINT ./server -c config.dev.yaml
