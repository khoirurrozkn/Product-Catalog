package controller

import (
	"net/http"
	"server/model"
	"server/model/dto/response"
	"server/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavoriteController struct {
	pu usecase.FavoriteUsecase
	rg *gin.RouterGroup
}

func (pc *FavoriteController) CreateHandler(c *gin.Context) {
	var newFavorite model.Favorite

	if err := c.ShouldBindJSON(&newFavorite); err != nil {
		response.SendSingleResponseError(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	data, err := pc.pu.CreateFavorite(newFavorite)
	if err != nil {
		response.SendSingleResponseError(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	response.SendSingleResponseCreated(
		c,
		data,
		"Success Create new favorite",
	)
}

func (pc *FavoriteController) getAllHandler(c *gin.Context) {
	order := c.DefaultQuery("order", "created_at")
	sort := c.DefaultQuery("sort", "DESC")
	limit := 1
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	if err != nil {
		response.SendSingleResponseError(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	validOrderBy := map[string]bool{
		"created_at": true,
	}

	validSort := map[string]bool{
		"ASC":  true,
		"DESC": true,
	}

	if valid := validOrderBy[order]; !valid {
		order = "created_at"
	}

	if valid := validSort[sort]; !valid {
		sort = "DESC"
	}

	favorite, paging, err := pc.pu.GetAllFavorite(
		order,
		sort,
		page,
		limit,
	)

	if err != nil {
		response.SendSingleResponseError(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	response.SendSinglePageResponse(
		c,
		favorite,
		"Success get list Favorite",
		paging,
	)
}

func (pc *FavoriteController) deleteByIdHandler(c *gin.Context) {
	favoriteId := c.Param("id")

	data, err := pc.pu.DeleteFavoriteById(favoriteId)
	if err != nil {
		response.SendSingleResponseError(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	response.SendSingleResponse(
		c,
		data,
		"Success deleted Favorite",
	)
}

func (pc *FavoriteController) Route() {
	router := pc.rg.Group("/favorite")
	router.POST("", pc.CreateHandler)
	router.GET("/", pc.getAllHandler)
	router.DELETE("/:id", pc.deleteByIdHandler)
}

func NewFavoriteController(pu usecase.FavoriteUsecase, routerGroup *gin.RouterGroup) *FavoriteController {
	return &FavoriteController{
		pu: pu,
		rg: routerGroup,
	}
}
