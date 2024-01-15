package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Paginate(ctx *gin.Context) (int, int, int) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return 0, 0, 0
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return 0, 0, 0
	}

	offset := (page - 1) * limit

	return page, limit, offset
}
