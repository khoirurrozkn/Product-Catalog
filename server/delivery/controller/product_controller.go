package controller

import (
	"net/http"
	"server/model"
	"server/model/dto/response"
	"server/usecase"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	uc usecase.ProductUsecase
	rg *gin.RouterGroup
}

func (cc *ProductController) CreateHandler(c *gin.Context) {
	var newProduct model.Product
	
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		response.SendSingleResponseError(
			c, 
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	data, err := cc.uc.CreateProduct(newProduct)
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
		"Success Create new Product",
	)
}

func (cc *ProductController) Route(){
	router := cc.rg.Group("/product")
	router.POST("", cc.CreateHandler)
}

func NewProductController(uc usecase.ProductUsecase, router *gin.Engine) *ProductController {
	return &ProductController{
		uc: uc,
		rg: &router.RouterGroup,
	}
}