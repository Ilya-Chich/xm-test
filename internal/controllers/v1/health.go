package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Build     string   `json:"build"`
	Status    int      `json:"status"`
	Messages  []string `json:"messages,omitempty"`
	Services  []string `json:"services"`
	Variables []string `json:"variables"`
}

const (
	healthQueryFormatNone       = "none"
	healthQueryFormatJSON       = "json"
	healthQueryFormatPrometheus = "prometheus"
)

func (c *CompanyController) Health(ctx *gin.Context) {
	format := ctx.Query("format")
	if format == "" {
		format = healthQueryFormatNone
	}
	switch format {
	case healthQueryFormatJSON:
		req := new(HealthResponse)
		req.Build = c.buildVersion
		req.Status = http.StatusOK
		ctx.JSON(http.StatusOK, req)
	case healthQueryFormatPrometheus:
		ctx.Writer.WriteHeader(http.StatusOK)
		payload := fmt.Sprintf("status %d", http.StatusOK)
		_, _ = ctx.Writer.WriteString(payload)
	}
}
