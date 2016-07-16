package server

import (
	"net"
	"log"
)

type ServerConfig struct {
	Host string
	Port string
	Type string
}

type Server struct {
	config ServerConfig
	listener func(net.Conn)
	server net.Listener
}

func NewServer(config ServerConfig, listener func(net.Conn)) *Server {
	s := Server{config: config, listener: listener}
	s.setup()
	return &s
}

func (this *Server) setup() {
	url := this.config.Host + ":" + this.config.Port
	serverListener, err := net.Listen(this.config.Type, url)
	if err != nil {
		log.Fatalln(err)
	}
	this.server = serverListener
}

func (this *Server) Start() {
	log.Println("Starting server on:", this.config.Host + ":" + this.config.Port)
	for {
		connection, err := this.server.Accept()
		if err != nil {
			log.Println(err)
		}

		go this.listener(connection)
	}
}

func (this *Server) GetPort() string {
	return this.config.Port
}

func (this *Server) GetHost() string {
	return this.config.Host
}

func (this *Server) GetType() string {
	return this.config.Type
}