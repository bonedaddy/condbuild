package main

import (
	"flag"

	"github.com/postables/condbuild/core"
)

var (
	mode = flag.String("mode", "", "set mode")
)

func init() {
	flag.Parse()
}

func main() {
	core.Linux()
}
