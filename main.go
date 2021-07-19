package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	command := flag.String("t", "", "sdf")
	args := flag.String("args", "", "Set the args. (separated by spaces)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-cmd <command>] [-args <the arguments (separated by spaces)>]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	println(*command)
	println(*args)

	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run("0.0.0.0:81")
}
