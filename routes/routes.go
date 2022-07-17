package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/maysapereira/go-api-rest-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.Run(":5000") //mudei a porta pois a 8080 não estava funcionando
}