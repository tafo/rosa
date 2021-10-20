package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/tafo/rosa/internal/auth/models"
	"net/http"
)

type AccountHandler struct {
	manager Manager
}

type Manager interface {
	Register(account models.Account) (AuthResponse, error)
}

func NewAccountHandler(manager Manager) AccountHandler {
	return AccountHandler{manager: manager}
}

func (ah AccountHandler) register(context *gin.Context) {
	var request RegisterRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.IndentedJSON(http.StatusBadRequest, context.Error(err))
		return
	}

	user := request.ToEntity()
	response, err := ah.manager.Register(user)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, context.Error(err))
		return
	}

	context.IndentedJSON(http.StatusCreated, response)
}
