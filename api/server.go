package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/thangnguyenvn98/simplebank/db/sqlc"
)

// Server serves HTTP requests for our banksing service

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// New server create new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)

	server.router = router
	return server
}

// Starts run the http srever on specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
