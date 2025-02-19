FROM golang:1.20-alpine

RUN apk update && apk add --no-cache \
    bash \
    curl \
    git \
    build-base

RUN mkdir -p /app
WORKDIR /app

COPY ./src /app/

RUN adduser -D -h /home/ctfuser -s /bin/bash ctfuser

RUN chown -R ctfuser:ctfuser /app && \
    chmod -R 550 /app

RUN chmod 770 /app/userEXE && \
    chmod 770 /app/userFile && \
    chown -R ctfuser:ctfuser /app/userEXE /app/userFile

USER root

RUN go build -o server .

USER ctfuser

EXPOSE 20000

CMD ["./server"]
