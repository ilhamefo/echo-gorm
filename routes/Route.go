package routes

import (
	"fmt"
	"gorm-echo/model"
	"gorm-echo/request"
	"gorm-echo/response"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) func(echo.Context) error {
	return func(ctx echo.Context) error {
		err := db.AutoMigrate(&model.User{})
		if err != nil {
			return response.ErrorResponse(err, ctx)
		}

		err = db.AutoMigrate(&model.Links{})
		if err != nil {
			return response.ErrorResponse(err, ctx)
		}

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "Migration successful",
		})
	}
}

func HelloWorld() func(echo.Context) error {
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"Hello": "World",
		})
	}
}

func GetAllUsers(db *gorm.DB) func(echo.Context) error {
	return func(ctx echo.Context) error {
		var users []response.User
		q := db.Omit("password").Find(&users)

		if q.Error != nil {
			return response.ErrorResponse(q.Error, ctx)
		}
		return ctx.JSON(http.StatusOK, users)
	}
}

func GetOneUser(db *gorm.DB) func(echo.Context) error {
	return func(ctx echo.Context) error {
		var user response.User

		id := ctx.Param("id")

		q := db.Where("id", id).Omit("password").First(&user)

		if q.Error != nil {
			return response.ErrorResponse(q.Error, ctx)
		}
		return ctx.JSON(http.StatusOK, user)
	}
}

func UpdateUser(db *gorm.DB) func(echo.Context) error {
	return func(ctx echo.Context) error {
		var user request.User
		id := ctx.Param("id")

		q := db.Where("id", id).First(&user)

		if q.Error != nil {
			return response.ErrorResponse(q.Error, ctx)
		}

		// binding data
		fmt.Println("ctx.Request().Body")
		fmt.Println(ctx.Request().Body)
		if err := ctx.Validate(ctx.Request().Body); err != nil {
			fmt.Println(err)
			return response.ErrorResponse(err, ctx)
		}

		err := ctx.Bind(&user)
		if err != nil {
			return response.ErrorResponse(q.Error, ctx)
		}

		q = q.Model(&user).Updates(&user)
		if q.Error != nil {
			return response.ErrorResponse(q.Error, ctx)
		}
		return ctx.JSON(http.StatusOK, user)
	}
}
