package api

import "github.com/nhutHao02/social-network-chat-service/internal/api/http"

type Server struct {
	// http server
	HTTPServer *http.HTTPServer
	// grpc ser ver

}

func NewSerVer(httpServer *http.HTTPServer) *Server {
	return &Server{HTTPServer: httpServer}
}
