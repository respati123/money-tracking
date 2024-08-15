package util

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/respati123/money-tracking/internal/constants"
	"github.com/respati123/money-tracking/internal/entity"
)

func GenerateNumber(digit int) int {
	if digit <= 0 {
		return 0
	}

	rand.Seed(time.Now().UnixNano())
	min := 1
	for i := 1; i < digit; i++ {
		min *= 10
	}
	max := min*10 - 1

	return rand.Intn(max-min+1) + min

}

func GetUserData(ctx *gin.Context) (entity.User, bool) {
	user := ctx.Value(constants.UserData)
	if user != nil {
		return user.(entity.User), true
	}
	return entity.User{}, false
}
