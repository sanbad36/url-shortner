package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sanbad36/url-shortner/api/routersetup"
	"github.com/stretchr/testify/assert"
)

func TestGetByShortID(t *testing.T) {
	router := gin.Default()
	routersetup.SetupRouters(router)

	req, err := http.NewRequest("GET", "/api/v1/abcd12", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
	expected := `{"error":"Data not found for given ShortID"}`
	assert.JSONEq(t, expected, resp.Body.String())
}