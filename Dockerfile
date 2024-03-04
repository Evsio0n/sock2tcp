FROM golang:1.21 as builder
#Copy from outside to inside

LABEL authors="evsio0n"


RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go mod tidy
RUN go build -o sock2tcp main.go
RUN cp sock2tcp /tmp/sock2tcp

FROM alpine:3.12 as runner
COPY --from=builder /tmp/sock2tcp /usr/local/bin/sock2tcp
RUN ls -la /usr/local/bin/sock2tcp
RUN chmod +x /usr/local/bin/sock2tcp
RUN mkdir /app
WORKDIR /app
ENV UNIX_SOCKET_PATH=""
ENV TCP_BIND_ADDRESS=""

CMD ["/usr/local/bin/sock2tcp"]


