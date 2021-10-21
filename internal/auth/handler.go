package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AccountHandler struct {

}

var Handler AccountHandler

func (ah AccountHandler) Register(context *gin.Context) {
	var request RegisterRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.IndentedJSON(http.StatusBadRequest, context.Error(err))
		return
	}

	account := request.ToEntity()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, context.Error(err))
		return
	}

	account.Password = string(hashedPassword)
	if err = Repo.Insert(&account); err != nil {
		context.IndentedJSON(http.StatusBadRequest, context.Error(err))
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    account.Id,
		"email": account.Email,
	})

	secret := []byte(viper.GetString("JWT_SECRET"))
	signedToken, err := token.SignedString(secret)

	context.IndentedJSON(http.StatusCreated, RegisterResponse{Token: signedToken})
}