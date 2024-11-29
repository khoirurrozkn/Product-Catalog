package controller

import (
	"net/http"
	"server/model"
	"server/model/dto"
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

func (cc *ProductController) getAllHandler(c *gin.Context) {
    product, err := cc.uc.GetProduct()
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
		product,
		"Success get list Product pagination",
		dto.Paging{},
	)
}

func (cc *ProductController) updateHandler(c *gin.Context){
	var updatedProduct model.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		response.SendSingleResponseError(
			c, 
			http.StatusBadRequest,
			err.Error(),
		)

        return
	}

	data, err := cc.uc.UpdateProductById(updatedProduct)
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
		"Success updated Product",
	)
}

func (cc *ProductController) deleteByIdHandler(c *gin.Context){
	idProduct := c.Param("id")

	err := cc.uc.DeleteProductById(idProduct)
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
		nil,
		"Success deleted Product",
	)
}

func (cc *ProductController) Route(){
	router := cc.rg.Group("/product")
	router.POST("", cc.CreateHandler)
	router.GET("", cc.getAllHandler)
	router.PUT("", cc.updateHandler)
	router.DELETE("/:id", cc.deleteByIdHandler)
}

func NewProductController(uc usecase.ProductUsecase, routerGroup *gin.RouterGroup) *ProductController {
	return &ProductController{
		uc: uc,
		rg: routerGroup,
	}
}