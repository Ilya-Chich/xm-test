package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"time"
	"xm-test-ilya-chicherin/internal/controllers/v1/view"
	"xm-test-ilya-chicherin/usecase"
)

type AuthController struct {
	AuthUC usecase.Auth
}

func NewAuthController(authUC usecase.Auth) *AuthController {
	return &AuthController{AuthUC: authUC}
}

func (a *AuthController) Login(ctx *gin.Context) {
	var req view.AuthRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	token, err := generateJWT(string(req.Email))
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": token,
		"token_type":   "Bearer",
		"expires_in":   3600,
	})
}
func generateJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": "1", // TODO create a real authorization
		"email":   email,
		"exp":     time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}
