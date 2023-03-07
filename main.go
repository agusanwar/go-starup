package main

import (
	"backend_startup/handler"
	"backend_startup/users"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/backend_startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// user
	userRepository := users.NewRepository(db)
	// services
	userServices := users.NewServices(userRepository)

	// routes register user
	userHandler := handler.NewUserHandler(userServices)

	router := gin.Default()
	api := router.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)
	router.Run()

	// router := gin.Default()
	// router.GET("/handler", handler)
	// router.Run()

}

// func handler(c *gin.Context) {
// 	dsn := "root:@tcp(127.0.0.1:3306)/backend_startup?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	// get data from database
// 	var users []users.User
// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users)
// }
