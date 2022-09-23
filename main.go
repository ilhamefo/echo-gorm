package main

import (
	"fmt"
	"gorm-echo/db"
	"gorm-echo/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func main() {
	db := db.ConnectDB()
	handleRequest(db)
}

type User struct {
	Id              string `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	EmailVerifiedAt string `json:"email_verified_at"`
	BacgroundColor  string `json:"background_color"`
	TextColor       string `json:"text_color"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
}

func getAllUsers(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var users []User
		db.Find(&users)
		// fmt.Println("{}", users)

		return c.JSON(http.StatusOK, users)
	}
}

func getOneUser(db *gorm.DB) func(echo.Context) error {
	return func(ctx echo.Context) error {
		var user []User
		id := ctx.Param("id")

		fmt.Println(id)
		db.Where("id", id).First(&user)
		return ctx.JSON(http.StatusOK, user)

	}
}

func handleRequest(db *gorm.DB) {

	migrate(db)
	e := echo.New()
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.GET("/users", getAllUsers(db))
	e.GET("/user/:id", getOneUser(db))
	e.Logger.Fatal(e.Start(":1323"))
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println(err)
	}

	err = db.AutoMigrate(&model.Links{})
	if err != nil {
		fmt.Println(err)
	}
}
