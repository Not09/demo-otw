package main

import (
	"github.com/not09/demo-otw/controller"

	"github.com/not09/demo-otw/entity"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())


	// Requirecustomer Routes

	r.GET("/requirecustomers", controller.ListRequirecustomers)

	r.GET("/requirecustomer/:id", controller.GetRequirecustomers)

	r.POST("/requirecustomers", controller.CreateRequirecustomer)

	r.PATCH("/requirecustomers", controller.UpdateRequirecustomer)

	r.DELETE("/requirecustomers/:id", controller.DeleteRequirecustomer)

	// Labheadfa Routes

	r.GET("/labheadfas", controller.ListLabheadfas)

	r.GET("/labheadfa/:id", controller.GetLabheadfas)

	r.POST("/labheadfas", controller.CreateLabheadfa)

	r.PATCH("/labheadfas", controller.UpdateLabheadfa)

	r.DELETE("/labheadfas/:id", controller.DeleteLabheadfa)

	// Engheadfa Routes

	r.GET("/engheadfas", controller.LisEngheadfas)

	r.GET("/engheadfa/:id", controller.GetEngheadfas)

	r.POST("/engheadfas", controller.CreateEngheadfa)

	r.PATCH("/engheadfas", controller.UpdateEngheadfa)

	r.DELETE("/engheadfas/:id", controller.DeleteEngheadfa)

	// Run the server

	r.Run()

}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

	  c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	  c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	  c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

	  c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")


	  if c.Request.Method == "OPTIONS" {

		c.AbortWithStatus(204)

		return

	  }


	  c.Next()

	}

  }