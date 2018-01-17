package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

  // Set the router as the default one provided by Gin
  router = gin.Default()

  // Initialize some process
  //InitKafkaConsumer()
  initializeRoutes()

  // Start serving the application
  router.Run("0.0.0.0:8000")

}
