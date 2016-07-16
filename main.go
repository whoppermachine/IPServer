package main

import (
	"log"
	"os"
	"net"
	"github.com/whoppermachine/IPServer/server"
	"fmt"
)

const HOST = "127.0.0.1"
const PORT = "1234"
const TYPE = "tcp"

const NEW_LINE = "\r\n"

func main() {
	log.SetOutput(os.Stdout)

	config := server.ServerConfig{
		Host: HOST,
		Port: PORT,
		Type: TYPE,
	}

	minServer := server.NewServer(config, handler)
	minServer.Start()
}

func handler(conn net.Conn) {
	address := conn.RemoteAddr().(*net.TCPAddr)
	log.Println("Got a request from:", address.IP)

	body := fmt.Sprint(address.IP)
	response := createHttpHeader(body)

	conn.Write(response)
}

func createHttpHeader(body string) []byte {
	bodyLen := len(body)

	header := "HTTP/1.1 200 OK" + NEW_LINE
	header += "Content-Type: text/plain" + NEW_LINE
	header += fmt.Sprint("Content-Length: ", bodyLen, NEW_LINE)
	header += "Connection: close" + NEW_LINE
	header += NEW_LINE
	header += body

	return []byte(header)
}