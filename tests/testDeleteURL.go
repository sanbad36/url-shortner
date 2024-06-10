package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sanbad36/url-shortner/api/routersetup"
	"github.com/stretchr/testify/assert"
)

func TestDeleteURL(t *testing.T) {
	router := gin.Default()
	routersetup.SetupRouters(router)

	req, err := http.NewRequest("DELETE", "/api/v1/abcd12", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	expected := `{"error":"Unable to delete the shortened Link"}`
	assert.JSONEq(t, expected, resp.Body.String())
}