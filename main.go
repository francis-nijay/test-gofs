package main

import (
	"context"
	"flag"
	"fmt"
	"os"
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

	fmt.Fprintf("Performin Bind og staging to target path")
	if err := gofsutil.BindMount(context.Background(), stagingPath, targetPath, mntFlags...); err != nil {
		return status.Error(codes.Internal, utils.GetMessageWithRunID(rid, "error publish volume to target path: %v", err))
	}

}
