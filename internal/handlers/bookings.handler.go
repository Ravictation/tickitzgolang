package handlers

import (
	"github.com/Ravictation/tickitzgolang/config"
	"github.com/Ravictation/tickitzgolang/internal/models"
	"github.com/Ravictation/tickitzgolang/internal/pkg"
	"github.com/Ravictation/tickitzgolang/internal/repositories"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type Handler_Bookings struct {
	*repositories.Repo_Bookings
}

func NewBookings(r *repositories.Repo_Bookings) *Handler_Bookings {
	return &Handler_Bookings{r}
}

func (h *Handler_Bookings) Get_Bookings(ctx *gin.Context) {
	var bookings models.Bookings

	page := ctx.Query("page")
	limit := ctx.Query("limit")
	time_schedule := ctx.Query("time_schedule")
	user := ""

	if err := ctx.ShouldBind(&bookings); err != nil {
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Get_Data(&bookings, page, limit, user, time_schedule)
	if err != nil {
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	pkg.NewRes(200, response).Send(ctx)
}

func (h *Handler_Bookings) Get_Bookings_by_User(ctx *gin.Context) {
	var bookings models.Bookings

	page := ctx.Query("page")
	limit := ctx.Query("limit")
	time_schedule := ctx.Query("time_schedule")
	user := ctx.MustGet("userId").(string)

	if err := ctx.ShouldBind(&bookings); err != nil {
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Get_Data(&bookings, page, limit, user, time_schedule)
	if err != nil {
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	pkg.NewRes(200, response).Send(ctx)
}

func (h *Handler_Bookings) Post_Bookings(ctx *gin.Context) {
	var bookings models.Bookingsset

	bookings.Id_user = ctx.MustGet("userId").(string)

	if err := ctx.Bind(&bookings); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	var err_val error
	_, err_val = govalidator.ValidateStruct(&bookings)
	if err_val != nil {
		pkg.NewRes(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}

	response, err := h.Insert_Data(&bookings)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.NewRes(200, &config.Result{Message: response}).Send(ctx)
}
