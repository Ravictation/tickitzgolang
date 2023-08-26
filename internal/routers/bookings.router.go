package routers

import (
	"github.com/Ravictation/tickitzgolang/internal/handlers"
	"github.com/Ravictation/tickitzgolang/internal/middleware"
	"github.com/Ravictation/tickitzgolang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func bookings_router(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/bookings")

	repo := repositories.NewBookings(d)
	handler := handlers.NewBookings(repo)

	route.GET("/user", middleware.Authjwt("user", "admin"), handler.Get_Bookings_by_User)
	route.GET("/", handler.Get_Bookings)
	route.POST("/", middleware.Authjwt("user", "admin"), handler.Post_Bookings)
}
