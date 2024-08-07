package converter

import "github.com/gin-gonic/gin"

type Converter struct {
	Ctx           *gin.Context
	UserConverter UserConverter
}
