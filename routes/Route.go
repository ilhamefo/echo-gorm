package routes

import (
	"gorm-echo/model"
	"gorm-echo/response"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) func(echo.Context) error {
	return func(ctx echo.Context) error {
		err := db.AutoMigrate(&model.User{})
		if err != nil {
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"message": "Migration error: " + err.Error(),
			})
		}

		err = db.AutoMigrate(&model.Links{})
		if err != nil {
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"message": "Migration error: " + err.Error(),
			})
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
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"message": "Error: " + q.Error.Error(),
			})
		}
		return ctx.JSON(http.StatusOK, users)
	}
}

func GetOneUser(db *gorm.DB) func(echo.Context) error {
	return func(ctx echo.Context) error {
		var user []response.User
		id := ctx.Param("id")

		q := db.Where("id", id).Omit("password").First(&user)

		if q.Error != nil {
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"message": "Error: " + q.Error.Error(),
			})
		}
		return ctx.JSON(http.StatusOK, user)
	}
}
