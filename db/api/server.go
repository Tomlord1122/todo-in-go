package api

import (
	db "github.com/Tomlord1122/todo-in-go/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	query *db.Queries
	route *gin.Engine
}

func NewServer(query *db.Queries) *Server {
	server := &Server{
		query: query,
	}
	router := gin.Default()

	// Add routes here

	router.GET("/todos", server.listTodos)
	router.POST("/todos", server.createTodo)
	router.GET("/todos/:id", server.getTodo)
	router.PUT("/todos/:id", server.updateTodo)
	router.DELETE("/todos/:id", server.deleteTodo)
	server.route = router
	return server
}

// errorResponse is a generic error response to return
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// Start runs the HTTP Server
func (server *Server) Start(address string) error {
	return server.route.Run(address)
}
