package router

import "github.com/gin-gonic/gin"

func Init() {
	r := gin.Default()

	api(r)

	err := r.Run(":8080")

	if err != nil {
		return
	}
}
