package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/sametdmr/simplebank/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	router.PATCH("/accounts", server.updateAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
