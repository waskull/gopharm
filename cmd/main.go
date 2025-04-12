package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	userHandler "github.com/waskull/gopharm/cmd/api/handlers/user"
	"github.com/waskull/gopharm/internal/repositories/mysql"
	userRepo "github.com/waskull/gopharm/internal/repositories/mysql/user"
	userService "github.com/waskull/gopharm/internal/services/user"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando el archivo .env")
	}

	ginEngine := gin.Default()
	ginEngine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hola",
		})
	})

	databaseParameters := mysql.ProvideDatabaseParameters()
	database := mysql.ProvideDatabase(databaseParameters)

	userRepo := &userRepo.Repository{
		DB: database,
	}

	userService := &userService.Service{
		Repo: userRepo,
	}

	userHandler := &userHandler.Handler{
		UserService: userService,
	}

	ginEngine.GET("/users/:id", userHandler.GetPlayer)
	ginEngine.POST("/users", userHandler.CreatePlayer)
	ginEngine.DELETE("/users/:id", userHandler.DeletePlayer)
	ginEngine.GET("/users", userHandler.GetUsers)
	ginEngine.PATCH("/users/:id", userHandler.GetPlayer)

	log.Fatal(ginEngine.Run(":8080"))
}
