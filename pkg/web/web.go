package web

type WebServer struct {
	port int
}

func (server WebServer) Start() {
	if server.port == 0 {
		server.port = 8080
	}
}