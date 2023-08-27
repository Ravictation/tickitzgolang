package handlers

import (
	"github.com/Ravictation/tickitzgolang/config"
	"github.com/Ravictation/tickitzgolang/internal/models"
	"github.com/Ravictation/tickitzgolang/internal/pkg"
	"github.com/Ravictation/tickitzgolang/internal/repositories"
	"github.com/gin-gonic/gin"
)

type Handler_Times_Schedules struct {
	*repositories.Repo_Times_Schedules
}

func NewTimesSchedules(r *repositories.Repo_Times_Schedules) *Handler_Times_Schedules {
	return &Handler_Times_Schedules{r}
}

func (h *Handler_Times_Schedules) Get_Times_Schedules(ctx *gin.Context) {
	var times_schedules models.Times_Scheduless

	page := ctx.Query("page")
	limit := ctx.Query("limit")
	movie := ctx.Query("movie")
	location_schedule := ctx.Query("location")
	time := ctx.Query("time")
	date := ctx.Query("date")

	if err := ctx.ShouldBind(&times_schedules); err != nil {
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Get_Data(&times_schedules, page, limit, location_schedule, time, date, movie)
	if err != nil {
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	pkg.NewRes(200, response).Send(ctx)
}

func (h *Handler_Times_Schedules) Get_Times_Schedules_by_Id(ctx *gin.Context) {
	var times_schedules models.Times_Scheduless

	times_schedules.Id_time_schedule = ctx.Param("id")
	if err := ctx.ShouldBind(&times_schedules); err != nil {
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Get_Data_by_Id(&times_schedules)
	if err != nil {
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	pkg.NewRes(200, response).Send(ctx)
}
