package main

import (
	"github.com/joho/godotenv"
	"io"
	"log"
	"net"
	"os"
)

var (
	unixSocketPath string
	tcpAddress     string
)

func init() {

	//read vars from env
	godotenv.Load()
	unixSocketPath = os.Getenv("UNIX_SOCKET_PATH")
	tcpAddress = os.Getenv("TCP_BIND_ADDRESS")
}

func main() {
	// 监听 TCP 地址
	listener, err := net.Listen("tcp", tcpAddress)
	if err != nil {
		log.Fatalf("Error listening on %s: %v", tcpAddress, err)
	}
	defer listener.Close()

	log.Printf("TCP server listening on %s, Fowarding from Unix socket %s", tcpAddress, unixSocketPath)

	// 接受连接
	for {
		tcpConn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting TCP connection: %v", err)
			continue
		}
		go handleTCPConnection(tcpConn)
	}
}

func handleTCPConnection(tcpConn net.Conn) {
	defer tcpConn.Close()

	// 连接到 Unix 域套接字
	unixConn, err := net.Dial("unix", unixSocketPath)
	if err != nil {
		log.Printf("Error connecting to Unix socket %s: %v", unixSocketPath, err)
		return
	}
	defer unixConn.Close()

	// 使用 io.Copy 双向传输数据
	go io.Copy(tcpConn, unixConn)
	io.Copy(unixConn, tcpConn)
}
