package handlers

import (
	"github.com/Ravictation/tickitzgolang/config"
	"github.com/Ravictation/tickitzgolang/internal/models"
	"github.com/Ravictation/tickitzgolang/internal/pkg"
	"github.com/Ravictation/tickitzgolang/internal/repositories"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type Handler_Movie struct {
	*repositories.Repo_Movies
}

func NewMovies(r *repositories.Repo_Movies) *Handler_Movie {
	return &Handler_Movie{r}
}

func (h *Handler_Movie) Get_Movies(ctx *gin.Context) {
	var movies models.Movies

	page := ctx.Query("page")
	limit := ctx.Query("limit")
	search := ctx.Query("search")
	orderby := ctx.Query("order_by")

	if err := ctx.ShouldBind(&movies); err != nil {
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Get_Data(&movies, page, limit, search, orderby)
	if err != nil {
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	pkg.NewRes(200, response).Send(ctx)
}
func (h *Handler_Movie) Post_Movies(ctx *gin.Context) {
	var moviesset models.Moviesset

	if err := ctx.Bind(&moviesset); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	moviesset.Image = ctx.MustGet("image").(string)
	moviesset.Cover_image = ctx.MustGet("image2").(string)

	var err_val error
	_, err_val = govalidator.ValidateStruct(&moviesset)
	if err_val != nil {
		pkg.NewRes(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}

	response, err := h.Insert_Data(&moviesset)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.NewRes(200, &config.Result{Message: response}).Send(ctx)
}
func (h *Handler_Movie) Put_Movies(ctx *gin.Context) {
	var moviesset models.Moviesset

	moviesset.Id_movie = ctx.Param("id")

	count_by_id := h.Get_Count_by_Id(moviesset.Id_movie)
	if count_by_id == 0 {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	moviesset.Image = ctx.MustGet("image").(string)
	moviesset.Cover_image = ctx.MustGet("image2").(string)

	if err := ctx.ShouldBind(&moviesset); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	var err_val error
	_, err_val = govalidator.ValidateStruct(&moviesset)
	if err_val != nil {
		pkg.NewRes(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}

	response, err := h.Update_Data(&moviesset)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.NewRes(200, &config.Result{Message: response}).Send(ctx)
}

func (h *Handler_Movie) Delete_Movies(ctx *gin.Context) {
	var movies models.Movies
	var movies_casts models.Movies_Casts
	var movies_genres models.Movies_Genres
	movies.Id_movie = ctx.Param("id")
	movies_casts.Id_movie = ctx.Param("id")
	movies_genres.Id_movie = ctx.Param("id")

	count_by_id := h.Get_Count_by_Id(movies.Id_movie)
	if count_by_id == 0 {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	if err := ctx.ShouldBind(&movies); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Delete_Data(&movies, &movies_casts, &movies_genres)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.NewRes(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.NewRes(200, &config.Result{Message: response}).Send(ctx)
}
