package utils

import (
	"fmt"
	"server/model/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ValidateRequestQueryParams(c *gin.Context) (dto.RequestQueryParam, error) {

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 0 {
		return dto.RequestQueryParam{}, fmt.Errorf("Invalid page number")
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if err != nil || limit <= 0 {
		return dto.RequestQueryParam{}, fmt.Errorf("Invalid limit number")
	}

	order := c.DefaultQuery("order", "name")
	sort := c.DefaultQuery("sort", "asc")

	return dto.RequestQueryParam{
		QueryParams: dto.QueryParams{
			Order: order,
			Sort: sort,
		},
		PaginationParam: dto.PaginationParam{
			Page: page,
			Limit: limit,
		},
	}, nil
}