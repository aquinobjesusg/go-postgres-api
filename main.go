package main

import (
	"go-postgres-api/config"
	"go-postgres-api/controllers"
	"go-postgres-api/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Inicializar la base de datos
	config.ConnectDatabase()

	// 2. Migrar los modelos
	config.DB.AutoMigrate(&models.User{})

	// 3. Configurar el router de Gin
	router := gin.Default()

	// 4. Configurar CORS personalizado
	router.Use(func(c *gin.Context) {
		// Establecer los encabezados CORS
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, access")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Content-Type", "application/json; charset=UTF-8")

		// Manejar solicitudes preflight OPTIONS
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// http://localhost:8080/users
	// http://localhost:8080/users/:id
	// 4. Definir las rutas de la API
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUserByID)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	// 5. Iniciar el servidor en el puerto 8080
	log.Println("Servidor iniciado en http://localhost:8080")
	router.Run(":8080")
}
