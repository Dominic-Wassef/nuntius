package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"

	"nuntius"
)

const version    = "0.9.9"

// main starts app
func main() {
	var filename = flag.String("config", "config.yml", "Config file location")
	flag.Parse()

	printBanner()

	nuntius.Start(*filename)
}

// print banner
func printBanner() {
	// print info
	fmt.Println("")
	log.Printf("******************************************")
	log.Printf("** %sNuntius%s v%s built in %s", "\033[31m", "\033[0m", version, runtime.Version())
	log.Printf("**----------------------------------------")
	log.Printf("** Running with %d Processors", runtime.NumCPU())
	log.Printf("** Running on %s", runtime.GOOS)
	log.Printf("******************************************")
}
