package handlers

import "github.com/gin-gonic/gin"

// Handler is an interface that defines a single method Do, which represents the main functionality of a handler.
// Handlers that implement this interface can be used to handle specific actions in a web application.
type Handler interface {
	// Do handles the main functionality of the handler based on the given gin.Context.
	Do(c *gin.Context)
}
