package routers

import (
	"github.com/Ravictation/tickitzgolang/internal/handlers"
	"github.com/Ravictation/tickitzgolang/internal/middleware"
	"github.com/Ravictation/tickitzgolang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func user(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	repo := repositories.NewUser(d)
	handler := handlers.NewUser(repo)

	route.POST("/", handler.PostData)
	route.PATCH("/:username", middleware.Authjwt("user", "admin"), middleware.UploadFile("image_user"), handler.UpdateData)
	route.GET("/", handler.GetAllData)
	route.GET("/:username", middleware.Authjwt("admin"), handler.GetDataUser)
	route.DELETE("/:username", middleware.Authjwt("admin", "user"), handler.DeleteData)

}
