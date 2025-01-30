package controller

import (
	"net/http"
	"server/middleware"
	"server/model"
	"server/model/dto/response"
	"server/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	pu usecase.ProductUsecase
	rg *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (pc *ProductController) CreateHandler(c *gin.Context) {
	var newProduct model.Product

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		response.SendSingleResponseError(
			c, 
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	data, err := pc.pu.CreateProduct(newProduct)
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

func (pc *ProductController) getAllHandler(c *gin.Context) {
	order := c.DefaultQuery("order", "created_at")
	sort := c.DefaultQuery("sort", "DESC")
	limit := 2
	page, err := strconv.Atoi( c.DefaultQuery("page", "1") )

	if err != nil {
		response.SendSingleResponseError(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

        return
    }

	validOrderBy := map[string]bool{
		"price":   true,
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

	product, paging, err := pc.pu.GetProduct(
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
		product,
		"Success get list Product pagination",
		paging,
	)
}

func (pc *ProductController) updateHandler(c *gin.Context){
	var updatedProduct model.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		response.SendSingleResponseError(
			c, 
			http.StatusBadRequest,
			err.Error(),
		)

        return
	}

	data, err := pc.pu.UpdateProductById(updatedProduct)
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

func (pc *ProductController) deleteByIdHandler(c *gin.Context){
	idProduct := c.Param("id")

	data, err := pc.pu.DeleteProductById(idProduct)
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
		"Success deleted Product",
	)
}

func (pc *ProductController) Route(){
	router := pc.rg.Group("/product")
	router.POST("", pc.authMiddleware.RequireToken("Admin"), pc.CreateHandler)
	router.GET("", pc.authMiddleware.RequireToken("User", "Admin"), pc.getAllHandler)
	router.PUT("", pc.authMiddleware.RequireToken("Admin"), pc.updateHandler)
	router.DELETE("/:id", pc.authMiddleware.RequireToken("Admin"), pc.deleteByIdHandler)
}

func NewProductController(pu usecase.ProductUsecase, routerGroup *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *ProductController {
	return &ProductController{
		pu: pu,
		rg: routerGroup,
	}
}