package api

import "github.com/gin-gonic/gin"

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func newErrorResponse(ctx *gin.Context, code int, err error) {
	ctx.AbortWithStatusJSON(code, gin.H{
		"error": ApiError{
			Code:    code,
			Message: err.Error(),
		},
	})
}
