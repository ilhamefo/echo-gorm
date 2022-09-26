package main

import (
	"fmt"
	"gorm-echo/db"
	"gorm-echo/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func main() {
	db := db.ConnectDB()
	handleRequest(db)
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	fmt.Println("validating....")
	fmt.Println(i)
	return cv.validator.Struct(i)
}

func handleRequest(db *gorm.DB) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.GET("/users", routes.GetAllUsers(db))
	e.GET("/", routes.HelloWorld())
	e.GET("/user/:id", routes.GetOneUser(db))
	e.PUT("/user/:id", routes.UpdateUser(db))
	e.GET("/migrate", routes.Migrate(db))
	e.Logger.Fatal(e.Start(":1323"))
}
