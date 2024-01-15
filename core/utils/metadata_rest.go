package utils

import "github.com/gin-gonic/gin"

func GenerateMetadata(total, page, limit, offset int) gin.H {
	return gin.H{
		"total":  total,
		"page":   page,
		"limit":  limit,
		"offset": offset,
	}
}
