package proxy

import (
	"log"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func Forward(target string, backendPath string) gin.HandlerFunc {
	remote, err := url.Parse(target)
	if err != nil {
		panic("invalid proxy target: " + err.Error())
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	return func(c *gin.Context) {
		log.Printf("Incoming → %s %s", c.Request.Method, c.Request.URL.Path)
		c.Request.Host = remote.Host
		c.Request.URL.Scheme = remote.Scheme
		c.Request.URL.Host = remote.Host
		// overwrite only the path
		c.Request.URL.Path = backendPath

		// log outgoing request
		log.Printf("Proxying → %s %s%s", c.Request.Method, target, backendPath)

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
