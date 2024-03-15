package handler

import (
	"github.com/gin-gonic/gin"
)

type OrderHandler interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	// Update(ctx *gin.Context)
	// Delete(ctx *gin.Context)
}
