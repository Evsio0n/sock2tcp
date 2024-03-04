# sock2tcp

sock2tcp is a simple tool to forward a  local socket to serve as a TCP server.
It is useful when you have a local socket  and you want to expose it to specific IP and port.

## Installation
```shell
go install github.com/evsio0n/sock2tcp
```


## Usage
1. use `.env` file to configure the server

| Parameter          | Description                   |
|--------------------|-------------------------------|
| `UNIX_SOCKET_PATH` | The path of the local socket  |
| `TCP_BIND_ADDRESS` | The port to expose the socket |

```dotenv
UNIX_SOCKET_PATH=/path/to/socket
TCP_BIND_ADDRESS=0.0.0.0:8080
```

2. use environment variables to configure the server

```shell
export UNIX_SOCKET_PATH=/path/to/socket
export TCP_BIND_ADDRESS=0.0.0.0:8080
./sock2tcp 
```


