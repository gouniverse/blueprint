package main

import (
	"os"

	"github.com/gouniverse/envenc"
)

func main() {
	envenc.NewCli().Run(os.Args)
}
