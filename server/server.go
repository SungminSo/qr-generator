package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

type Server struct {
	Bind 	string
	router 	*gin.Engine
}

func NewServer(bind string) *Server {
	r := gin.Default()

	r.Use(gin.Logger(), gin.Recovery(), cors.Default())
	registerRoutes(r)

	return &Server{
		Bind:   bind,
		router: r,
	}
}

func (s *Server) Run() error {
	return s.router.Run(s.Bind)
}