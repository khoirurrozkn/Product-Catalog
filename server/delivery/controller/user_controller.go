package controller

import (
	"net/http"
	"server/model"
	"server/model/dto/response"
	"server/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	pu usecase.UserUsecase
	rg *gin.RouterGroup
}

func (pc *UserController) CreateHandler(c *gin.Context) {
	var newUser model.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		response.SendSingleResponseError(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	data, err := pc.pu.CreateUser(newUser)
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
		"Success Create new User",
	)
}

func (pc *UserController) getAllHandler(c *gin.Context) {
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
		"updated_at": true,
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

	user, paging, err := pc.pu.GetAllUser(
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
		user,
		"Success get list User pagination",
		paging,
	)
}

func (pc *UserController) deleteByIdHandler(c *gin.Context) {
	idUser := c.Param("id")

	data, err := pc.pu.DeleteUserById(idUser)
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
		"Success deleted User",
	)
}

func (pc *UserController) Route() {
	router := pc.rg.Group("/user")
	router.POST("", pc.CreateHandler)
	router.GET("", pc.getAllHandler)
	router.DELETE("/:id", pc.deleteByIdHandler)
}

func NewUserController(pu usecase.UserUsecase, routerGroup *gin.RouterGroup) *UserController {
	return &UserController{
		pu: pu,
		rg: routerGroup,
	}
}
