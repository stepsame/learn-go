package test

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"learn-gin/initRouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserSave(t *testing.T) {
	username := "lisi"
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "user "+username+" has saved.", w.Body.String())
}

func TestUserSaveQuery(t *testing.T) {
	username := "lisi"
	age := "20"
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/user?name=%s", username), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf("user: %s, age: %s has saved.", username, age), w.Body.String())
}
