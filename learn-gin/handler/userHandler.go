package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserSave(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "user "+username+" has saved.")
}

func UserSaveByQuery(context *gin.Context) {
	username := context.Query("name")
	age := context.DefaultQuery("age", "20")
	context.String(http.StatusOK, fmt.Sprintf("user: %s, age: %s has saved.", username, age))
}
