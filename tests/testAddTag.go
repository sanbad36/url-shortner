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

func TestAddTag(t *testing.T) {
	router := gin.Default()
	routersetup.SetupRouters(router)

	payload := `{"shortID": "abcd12", "tag": "example"}`
	req, err := http.NewRequest("POST", "/api/v1/addTag", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
	expected := `{"error":"Data not found for the given ShortID"}`
	assert.JSONEq(t, expected, resp.Body.String())
}