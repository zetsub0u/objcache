package main

import (
	"fmt"
	"os"

	"github.com/zetsub0u/objcache/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
