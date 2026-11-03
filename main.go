package main

import (
	"log"
	"os"

	"github.com/elastic/go-sysinfo"
	"github.com/elastic/go-sysinfo/types"
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

func DumpGO() map[string]any {
	body := gin.H{}

	body["arch"] = sysinfo.Go().Arch
	body["maxprocs"] = sysinfo.Go().MaxProcs
	body["os"] = sysinfo.Go().OS
	body["version"] = sysinfo.Go().Version

	return body
}

func DumpProcesses() []types.Process {

	proc, err := sysinfo.Processes()
	if err != nil {
		return proc
	} else {
		return nil
	}
}

func handle(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"host":      HostDump(ctx),
		"go":        DumpGO(),
		"processes": DumpProcesses(),
	})
}

func main() {
	router := gin.Default()

	auth := router.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":    "bar1",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	auth.GET("/", handle)
	gin.SetMode(gin.ReleaseMode)

	var err error = nil
	if len(os.Args) >= 2 {
		err = router.Run(os.Args[1])
	} else {
		err = router.Run()
	}
	if err != nil {
		log.Fatal(err)
	}

}
