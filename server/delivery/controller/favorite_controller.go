package controller

import (
	"net/http"
	"server/middleware"
	"server/model/dto/request"
	"server/model/dto/response"
	"server/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type FavoriteController struct {
	pu usecase.FavoriteUsecase
	rg *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (pc *FavoriteController) CreateHandler(c *gin.Context) {
	var productId request.Favorite
	claims := c.MustGet("claims").(jwt.MapClaims)
    userId := claims["user_id"].(string)

	if err := c.ShouldBindJSON(&productId); err != nil {
		response.SendSingleResponseError(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	data, err := pc.pu.CreateFavorite(productId, userId)
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

func (pc *FavoriteController) getAllByIdUserHandler(c *gin.Context) {
	order := c.DefaultQuery("order", "created_at")
	sort := c.DefaultQuery("sort", "DESC")
	limit := 10
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	claims := c.MustGet("claims").(jwt.MapClaims)
    userId := claims["user_id"].(string)

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
		userId,
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
	claims := c.MustGet("claims").(jwt.MapClaims)
    userId := claims["user_id"].(string)

	data, err := pc.pu.DeleteFavoriteById(userId, favoriteId)
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
	router.POST("/", pc.authMiddleware.RequireToken("User"), pc.CreateHandler)
	router.GET("/", pc.authMiddleware.RequireToken("User"), pc.getAllByIdUserHandler)
	router.DELETE("/:id", pc.authMiddleware.RequireToken("User"), pc.deleteByIdHandler)
}

func NewFavoriteController(pu usecase.FavoriteUsecase, routerGroup *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *FavoriteController {
	return &FavoriteController{
		pu: pu,
		rg: routerGroup,
		authMiddleware: authMiddleware,
	}
}
