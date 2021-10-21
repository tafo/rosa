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
		context.IndentedJSON(http.StatusInternalServerError, context.Error(err))
		return
	}

	signedToken, err := generateToken(account)

	context.IndentedJSON(http.StatusCreated, RegisterResponse{Token: signedToken})
}

func (ah AccountHandler) Login(context *gin.Context) {
	var request LoginRequest
	if err := context.ShouldBind(&request); err != nil {
		context.IndentedJSON(http.StatusBadRequest, context.Error(err))
		return
	}

	account := request.ToEntity()
	result := Repo.db.First(&account, "email = ?", account.Email)
	if result.RowsAffected == 0 {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "invalid email address"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(request.Password)); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "invalid password"})
		return
	}

	signedToken, err := generateToken(account)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message" : "internal server error"})
		return
	}

	context.IndentedJSON(http.StatusOK, LoginResponse{Token: signedToken})
}

func generateToken(account Account) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    account.Id,
		"email": account.Email,
	})

	secret := []byte(viper.GetString("JWT_SECRET"))
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}