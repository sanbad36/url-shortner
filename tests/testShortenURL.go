package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sanbad36/url-shortner/api/routersetup"
	"github.com/stretchr/testify/assert"
)

func TestShortenURL(t *testing.T) {
	router := gin.Default()
	routersetup.SetupRouters(router)

	payload := `{"url": "https://www.facebook.com"}`
	req, err := http.NewRequest("POST", "/api/v1", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	expected := `{"expiry":0,"rate_limit":10,"rate_limit_reset":30,"short":"http://localhost:3000/","url":"https://www.facebook.com"}`
	assert.JSONEq(t, expected, resp.Body.String())
}