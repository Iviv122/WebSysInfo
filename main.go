package main

import (
	"github.com/elastic/go-sysinfo"
	"github.com/gin-gonic/gin"
)

var fqdn string = ""

func HostDump(ctx *gin.Context) map[string]any {
	host, err := sysinfo.Host()
	if err != nil {
		return nil
	}

	body := gin.H{}
	body["host"] = host.Info()
	body["uptime"] = host.Info().Uptime()
	mem, err := host.Memory()
	if err != nil {
		body["memory"] = nil
	} else {

		body["memory"] = mem
	}
	cpu, err := host.CPUTime()
	if err != nil {
		body["cpu"] = nil
	} else {
		body["cpu"] = cpu
	}
	return body
}

func dump() map[string]any {
	sysinfo.Go()
}

func handle(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"body": gin.H{
			"host": HostDump(ctx),
			"go":   dump(),
		},
	})
}

func main() {

	r := gin.Default()
	r.NoRoute(handle)
	r.NoMethod(handle)
	gin.SetMode(gin.ReleaseMode)
	r.Run()

}
