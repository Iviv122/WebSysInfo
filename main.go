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

func Dump() []types.Process {

	proc, err := sysinfo.Processes()
	if err != nil {
		return proc
	} else {
		return nil
	}
}

func handle(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"body": gin.H{
			"host":      HostDump(ctx),
			"go":        DumpGO(),
			"processes": Dump(),
		},
	})
}

func main() {
	r := gin.Default()
	r.NoRoute(handle)
	r.NoMethod(handle)
	gin.SetMode(gin.ReleaseMode)
	var err error = nil
	if len(os.Args) >= 2 {
		err = r.Run(os.Args[1])
	} else {
		err = r.Run()
	}
	if err != nil {
		log.Fatal(err)
	}
}
