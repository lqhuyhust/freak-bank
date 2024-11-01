package api

import (
	"freak-bank/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Server serves all HTTP request
type Server struct {
	store  *sqlc.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *sqlc.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts", server.listAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.PUT("/accounts/:id", server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	// add routes to router
	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"erorr": err.Error()}
}
