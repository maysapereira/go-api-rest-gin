package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/maysapereira/go-api-rest-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.Run(":5000")
}