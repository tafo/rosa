package todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ItemHandler struct {
}

var Handler ItemHandler

func (ih ItemHandler) GetAll(context *gin.Context) {
	var items []Item

	if err := Repo.GetAll(&items); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, items)
}

func (ih ItemHandler) Create(context *gin.Context) {
	var request AddRequest
	err := context.ShouldBind(&request)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": context.Error(err)})
		return
	}

	item := request.ToEntity()

	if err = Repo.Insert(&item); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, context.Error(err))
		return
	}

	context.IndentedJSON(http.StatusCreated, AddResponse{item})
}

func (ih ItemHandler) Complete(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	item := &Item{Id: id}
	if err := Repo.Complete(item); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, context.Error(err))
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message" : "Item is completed"})
}