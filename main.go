package main

import (
	"github.com/elastic/go-sysinfo"
)

func main() {
	process, err := sysinfo.Host()
	if err != nil {
		panic(err)
	}
	print(process.Info().Architecture)
}
