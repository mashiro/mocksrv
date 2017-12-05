package main

import "github.com/gin-gonic/gin"

// Responder respond to the request
type Responder interface {
	Respond(c *gin.Context)
}
