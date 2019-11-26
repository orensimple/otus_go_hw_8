package main

import (
	"log"

	"github.com/orensimple/otus_hw1_8/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
