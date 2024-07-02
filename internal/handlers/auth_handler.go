package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/respati123/money-tracking/internal/repository"
	"github.com/respati123/money-tracking/internal/service"
	"github.com/respati123/money-tracking/util"
	"gorm.io/gorm"
)

type AuthHandlerImpl struct {
	Service service.AuthService
}

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewAuthHandler(router *gin.RouterGroup, db gorm.DB, svc util.Config) {
	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	handler := &AuthHandlerImpl{Service: authService}
	router.GET("/create-user", handler.CreateUser)

}

// @Summary Create a new user
// @Description Create a new user with input payload
// @Tags hello
// @Accept  json
// @Produce  json
// @Success 200 {object} User
// @Router /create-user [get]
func (handler *AuthHandlerImpl) CreateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": "ok",
	})
}
