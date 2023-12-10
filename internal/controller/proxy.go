package controller

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Proxy(c *gin.Context) {
	d, _ := base64.StdEncoding.DecodeString("aHR0cDovL21vb2MyLWFucy5jaGFveGluZy5jb20=")
	target, _ := url.Parse(string(d))

	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.Director = func(req *http.Request) {
		req.Host = target.Host
		req.Body = c.Request.Body
		req.Method = c.Request.Method
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.Header.Set("Cookie", "UID=1")
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
