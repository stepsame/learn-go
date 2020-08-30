package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": fmt.Sprintf("hello gin %s method", strings.ToLower(context.Request.Method)),
	})
}
