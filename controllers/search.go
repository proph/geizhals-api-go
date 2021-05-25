package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proph/geizhals-api-go/helpers"
	"github.com/proph/geizhals-api-go/services"
)

func InitSearchController(router *gin.RouterGroup) {
	r := router.Group("/search")
	r.GET("/:query", getSearchResult())
	r.GET("/category/:category/:query", getSearchResult())
}

func getSearchResult() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		query := ctx.Param("query")
		category, err := convertParamToInt(ctx.Param("category"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(http.StatusBadRequest, err))
		} else {
			result := services.Search(query, category)
			ctx.JSON(http.StatusOK, result)
		}
	}
}

func convertParamToInt(p string) (int, error) {
	i, err := strconv.Atoi(p)
	if err != nil {
		return 0, err
	}
	return i, nil
}
