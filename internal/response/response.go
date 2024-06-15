package response

import (
	"github.com/gin-gonic/gin"
)

func RespondWithError(c *gin.Context, code int, errResponse ResponseError) {
	c.JSON(code, errResponse)
}

func RespondWithJSON(c *gin.Context, code int, payload interface{}) {
	c.JSON(code, payload)
}
