package response

import (
	"server/model/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSingleResponseCreated(c *gin.Context, data any, descriptionMsg string) {
	c.JSON(http.StatusCreated, 
		&SingleResponse{
			Status: Status{
				Code: http.StatusCreated,
				Description: descriptionMsg,
			},
			Data: data,
		},
	)
}

func SendSingleResponse(c *gin.Context, data any, descriptionMsg string) {
	c.JSON(http.StatusOK, 
		&SingleResponse{
			Status: Status{
				Code: http.StatusOK,
				Description: descriptionMsg,
			},
			Data: data,
		},
	)
}

func SendSinglePageResponse(c *gin.Context, data []any, descriptionMsg string, paging dto.Paging) {
	c.JSON(http.StatusOK, 
		&PagedResponse{
			Status: Status{
				Code: http.StatusOK,
				Description: descriptionMsg,
			},
			Data: data,
			Paging: paging,
		},
	)
}

func SendSingleResponseError(c *gin.Context, code int, errorMessage string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, 
		&Status{
			Code: code,
			Description: errorMessage,
		},
	)
}