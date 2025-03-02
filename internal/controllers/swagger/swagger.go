package swagger

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/swaggest/swgui/v4cdn"
)

func Register(r gin.IRouter, swFn func() (*openapi3.T, error), basePath string) error {
	swDoc, err := swFn()
	if err != nil {
		return err
	}
	var (
		once           sync.Once
		swaggerHandler = v4cdn.NewHandler("API "+basePath, basePath+"/docs.json", "/")
	)

	r.GET("/documentation/*any", gin.WrapH(swaggerHandler))
	r.GET("/docs.json", func(c *gin.Context) {
		once.Do(func() {
			if os.Getenv("DEBUG") == "true" {
				u, _ := url.Parse(c.Request.Header.Get("Referer"))
				uri := fmt.Sprintf("%s%s", os.Getenv("HTTP_OR_HTTPS"), os.Getenv("HTTP_HOST"))
				if os.Getenv("HTTP_PORT") != "80" && os.Getenv("HTTP_PORT") != "443" {
					uri = fmt.Sprintf("%s%s:%s", os.Getenv("HTTP_OR_HTTPS"), u.Hostname(), os.Getenv("HTTP_PORT"))
				}
				swDoc.Servers = append(
					[]*openapi3.Server{
						{
							URL: uri,
						},
					},
					swDoc.Servers[0:]...)
			}
		})
		c.JSON(http.StatusOK, &swDoc)
	})
	return nil
}
