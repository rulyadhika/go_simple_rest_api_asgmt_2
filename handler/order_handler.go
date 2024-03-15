package handler

import (
	"github.com/gin-gonic/gin"
)

type OrderHandler interface {
	Create(ctx *gin.Context)
	// FindOne(ctx *gin.Context)
	// Update(ctx *gin.Context)
	// Delete(ctx *gin.Context)
}
