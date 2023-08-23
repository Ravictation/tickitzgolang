package handlers

import (
	"github.com/Ravictation/tickitzgolang/config"
	"github.com/Ravictation/tickitzgolang/internal/models"
	"github.com/Ravictation/tickitzgolang/internal/pkg"
	"github.com/Ravictation/tickitzgolang/internal/repositories"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type Handler_Casts struct {
	*repositories.Repo_Casts
}

func NewCasts(r *repositories.Repo_Casts) *Handler_Casts {
	return &Handler_Casts{r}
}

func (h *Handler_Casts) Get_Casts(ctx *gin.Context) {
	var casts models.Casts

	page := ctx.Query("page")
	limit := ctx.Query("limit")

	if err := ctx.ShouldBind(&casts); err != nil {
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Get_Data(&casts, page, limit)
	if err != nil {
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	pkg.NewRes(200, response).Send(ctx)
}

func (h *Handler_Casts) Post_Casts(ctx *gin.Context) {
	var casts models.Casts

	if err := ctx.ShouldBind(&casts); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	var err_val error
	_, err_val = govalidator.ValidateStruct(&casts)
	if err_val != nil {
		pkg.NewRes(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}
	response, err := h.Insert_Data(&casts)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.NewRes(200, &config.Result{Message: response}).Send(ctx)
}
func (h *Handler_Casts) Put_Casts(ctx *gin.Context) {
	var casts models.Casts
	casts.Id_cast = ctx.Param("id")

	count_by_id := h.Get_Count_by_Id(casts.Id_cast)
	if count_by_id == 0 {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	if err := ctx.ShouldBind(&casts); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	var err_val error
	_, err_val = govalidator.ValidateStruct(&casts)
	if err_val != nil {
		pkg.NewRes(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}
	response, err := h.Update_Data(&casts)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.NewRes(200, &config.Result{Message: response}).Send(ctx)
}

func (h *Handler_Casts) Delete_Casts(ctx *gin.Context) {
	var casts models.Casts
	casts.Id_cast = ctx.Param("id")

	count_by_id := h.Get_Count_by_Id(casts.Id_cast)
	if count_by_id == 0 {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	if err := ctx.ShouldBind(&casts); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Delete_Data(&casts)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.NewRes(200, &config.Result{Message: response}).Send(ctx)
}
