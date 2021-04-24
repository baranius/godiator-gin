package healthcheck

import (
	"github.com/baranius/godiator"
	"github.com/baranius/godiator-echo/resources/healthcheck/healthcheck_handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ResourceTestSuite struct {
	suite.Suite
}

func TestResourceSuite(t *testing.T){
	suite.Run(t, new(ResourceTestSuite))
}

func (s *ResourceTestSuite) Test_HealthCheck() {
	// Given
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, ctx := gin.CreateTestContext(w)

	httpRequest := httptest.NewRequest(http.MethodGet, "/health-check", nil)

	mockGodiator := godiator.MockGodiator{}
	mockGodiator.OnSend = func(request interface{}, params ...interface{}) (interface{}, error) {
		return &healthcheck_handlers.HealthCheckResponse{
			RequestMethod: "GET",
			Success:       true,
		}, nil
	}

	resource := Resource{gdtr: &mockGodiator}
	ctx.GET("/health-check", resource.GetHealthCheck)

	// When
	ctx.ServeHTTP(w, httpRequest)

	// Then
	assert.Equal(s.T(), http.StatusOK, w.Result().StatusCode)
	assert.Equal(s.T(), "{\"request-method\":\"GET\",\"success\":true}", w.Body.String())
}
