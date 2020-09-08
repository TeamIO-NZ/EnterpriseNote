package web

import "go.uber.org/zap"

type WebServer struct {
	port int
}

func (server WebServer) Start() {
	if server.port == 0 {
		server.port = 8080
		zap.S().Warn("No webserver port config detected, using 8080.")
	}
}
