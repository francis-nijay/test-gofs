package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/dell/gofsutil"
)

func main() {
	stagingPath := flag.String("staging-path", "", "staging path")
	targetPath := flag.String("target-path", "", "staging path")

	flag.Parse()

	if *stagingPath == "" {
		fmt.Fprintf(os.Stderr, "staging-path argument is mandatory")
		os.Exit(1)
	}

	if *targetPath == "" {
		fmt.Fprintf(os.Stderr, "target-path argument is mandatory")
		os.Exit(1)
	}

	var mntFlags []string

	fmt.Printf("Performin Bind of staging to target path")
	if err := gofsutil.BindMount(context.Background(), *stagingPath, *targetPath, mntFlags...); err != nil {
		fmt.Printf("Error during Bind mount:", err)
	}

}
