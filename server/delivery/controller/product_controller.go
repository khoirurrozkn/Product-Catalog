package controller

import (
	"server/usecase"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	uc usecase.ProductUsecase
	rg *gin.RouterGroup
}

func (cc *ProductController) Route(){

}

func NewProductController(uc usecase.ProductUsecase, router *gin.Engine) *ProductController {
	return &ProductController{
		uc: uc,
		rg: &router.RouterGroup,
	}
}