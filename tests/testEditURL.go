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

func TestEditURL(t *testing.T) {
	router := gin.Default()
	routersetup.SetupRouters(router)

	payload := `{"url": "https://www.udemy.com", "expiry": 24}`
	req, err := http.NewRequest("PUT", "/api/v1/abcd12", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	expected := `{"error":"ShortID doesn't exists"}`
	assert.JSONEq(t, expected, resp.Body.String())
}
