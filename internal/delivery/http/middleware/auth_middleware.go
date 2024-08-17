package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/respati123/money-tracking/internal/constants"
	"github.com/respati123/money-tracking/internal/entity"
	"github.com/respati123/money-tracking/internal/util"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func AuthMiddleware(redis *redis.Client, viper *viper.Viper, log *zap.Logger, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		authorization := c.GetHeader("Authorization")
		token := strings.Split(authorization, "Bearer ")

		value, err := util.VerifiJwtToken(token[1], viper.GetString("JWT_SECRET_KEY"))

		if err != nil {
			log.Info("when verify jwt token", zap.Error(err))
			util.SendErrorResponse(c, http.StatusUnauthorized, "unauthorized", err)
			return
		}
		var uuidString string
		for key, val := range value {
			if key == "Payload" {
				uuidString = val.(string)
			}
		}

		key := fmt.Sprintf("%s_%s", uuidString, constants.Token)

		_, err = redis.Get(c, key).Result()

		if err != nil {
			log.Info("when verify jwt token to redis", zap.Error(err))
			util.SendErrorResponse(c, http.StatusUnauthorized, "unauthorized", err)
			return
		}

		var user entity.User
		err = db.Table("users").Where("uuid =? ", uuidString).First(&user).Error

		if err != nil {
			log.Info("when verify jwt token", zap.Error(err))
			util.SendErrorResponse(c, http.StatusUnauthorized, "unauthorized", err)
			return
		}

		c.Set(constants.UserData, user)
		c.Next()

	}
}
