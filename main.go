package main

import (
	"gorm-echo/db"
	"gorm-echo/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func main() {
	db := db.ConnectDB()
	handleRequest(db)
}

func handleRequest(db *gorm.DB) {
	e := echo.New()
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.GET("/users", routes.GetAllUsers(db))
	e.GET("/", routes.HelloWorld())
	e.GET("/user/:id", routes.GetOneUser(db))
	e.GET("/migrate", routes.Migrate(db))
	e.Logger.Fatal(e.Start(":1323"))
}
