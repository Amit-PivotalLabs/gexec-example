package main

import (
	"flag"
	"os"
)

func main() {
	exitCode := flag.Int("exitCode", 0, "exit code for this program to return")
	flag.Parse()
	os.Exit(*exitCode)
}
