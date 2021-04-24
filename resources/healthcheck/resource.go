package healthcheck

import (
	"github.com/baranius/godiator"
	"github.com/baranius/godiator-echo/resources/healthcheck/healthcheck_handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resource struct {
	gdtr godiator.IGodiator
}

func NewResource() Resource {
	return Resource{gdtr: godiator.GetInstance()}
}

func (h *Resource) GetHealthCheck(c *gin.Context) {
	req := &healthcheck_handlers.HealthCheckRequest{}
	response, err := h.gdtr.Send(req, c)

	if err != nil {
		c.String(http.StatusBadRequest, "failed")
		return
	}

	c.JSON(http.StatusOK, response)
}