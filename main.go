package main

import (
	"github.com/elastic/go-sysinfo"
	"github.com/gin-gonic/gin"
)

func HostDump(ctx *gin.Context) map[string]any {
	host, err := sysinfo.Host()
	if err != nil {
		return nil
	}
	body := gin.H{}
	body["host"] = host.Info()
	body["uptime"] = host.Info().Uptime()
	return body
}

func handle(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"body": HostDump(ctx),
	})
}

func main() {

	r := gin.Default()
	r.NoRoute(handle)
	r.NoMethod(handle)
	gin.SetMode(gin.ReleaseMode)
	r.Run()

}
