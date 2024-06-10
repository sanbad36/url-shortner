package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sanbad36/url-shortner/api/routersetup"
	"github.com/stretchr/testify/assert"
)

func TestGetTopDomains(t *testing.T) {
	router := gin.Default()
	routersetup.SetupRouters(router)

	req, err := http.NewRequest("GET", "/api/v1/top-domains", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	// Assuming response format {"udemy.com":6,"facebook.com":4}
	expected := `{"facebook.com":4,"udemy.com":6}`
	assert.JSONEq(t, expected, resp.Body.String())
}