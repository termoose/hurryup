package main

import (
	"flag"
	"hurryup/internal/display"
	"hurryup/internal/speed"
)

var backend = flag.String("backend", "ookla", "specify the speedtest backend to use [ookla]")
var output = flag.String("output", "anim", "select the output format of the speedtest [anim, json, promfile, prometheus]")

func main() {
	_, _ = speed.NewNetflix()
	return
	flag.Parse()
	var s speed.Tester

	switch *backend {
	default:
		s, _ = speed.NewOokla()
	}

	switch *output {
	case "json":
		_ = display.JSON(s)
	case "anim":
		_ = display.BubbleTea(s)
	case "promfile":
		_ = display.PromFile(s)
	case "prometheus":
		_ = display.Prometheus(s)
	default:
		_ = display.BubbleTea(s)
	}

}
