package main

import (
	"fmt"
	"mx-bridge/mxnet"
	"os"
)

func main() {
	fmt.Println("Starting MX-Bridge")

	args := os.Args
	br := mxnet.NewBridge()
	if len(args) > 1 {
		if args[1] == "-t" {
			br.Show = true
		}
	}
	br.Start()

}
