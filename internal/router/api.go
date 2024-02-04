package router

import (
	"github.com/crud/internal/handler"
	"github.com/gin-gonic/gin"
)

func api(r *gin.Engine) {
	router := r.Group("/api")

	{
		router.GET("/users", handler.GetUsers)
		router.GET("/users/:id", handler.ShowUsers)
		router.POST("/users", handler.PostUsers)
		router.PUT("/users", handler.PutUsers)
		router.DELETE("/users", handler.DeleteUsers)
	}
}
