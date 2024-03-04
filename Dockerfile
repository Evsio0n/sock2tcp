FROM golang:1.21 as builder
#Copy from outside to inside

LABEL authors="evsio0n"


RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go mod tidy
RUN go build -o sock2tcp main.go
RUN cp sock2tcp /tmp/sock2tcp

FROM alpine:3.12
COPY --from=builder /tmp/sock2tcp /usr/local/bin/sock2tcp

ENV UNIX_SOCKET_PATH=""
ENV TCP_BIND_ADDRESS=""
ENTRYPOINT ["/usr/local/bin/sock2tcp"]

