package server

import (
	// "goback/db"
	"goback/controller"
	"goback/logger"
	mongo "goback/mongodb"

	// "goback/pkg/scripts"
	"github.com/gin-gonic/gin"
)

func Init() {
	logger.Init()
	mongo.Init()

	// gin server
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		controller.Create()
		c.JSON(200, gin.H{
			"ping": "pong",
		})
	})

	r.Run()

	// es.Init()
	// db.Init()
	// redis.Init()
	// scriptsModule := scripts.NewScriptsModuleSingleton(db.Factory)
	// scriptsModule.GetScript().Init()
	// r := NewRouter()
	// r.Run(":" + "4000")

}
