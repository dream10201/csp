package main

import (
	"flag"
	"fmt"
	"github.com/csp/v2/route"
)

func main() {
	args := flag.String("port", "1516", "Set the args.")
	flag.Parse()
	route.Run(fmt.Sprintf(":%s", *args))
}
