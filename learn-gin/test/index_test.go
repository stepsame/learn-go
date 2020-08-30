package test

import (
	"github.com/go-playground/assert/v2"
	"learn-gin/initRouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHtml(t *testing.T) {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	//assert.Equal(t, "hello gin get method", w.Body.String())
}
