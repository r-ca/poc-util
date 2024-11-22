package main

import (
	"flag"
	"fmt"
	"os"
)

func parseCLIArgs() (string, string, string, string) {
	// コマンドライン引数の解析
	name := flag.String("name", "", "Repository name (required)")
	license := flag.String("license", "Unlicense", "License type: MIT or Unlicense (default: The Unlicense)")
	remoteURL := flag.String("remote", "", "Custom remote URL (optional, for debugging)")
	dir := flag.String("dir", ".", "Target directory for repository (default: current directory)")
	flag.Parse()

	if *name == "" {
		fmt.Println("Error: --name is required")
		os.Exit(1)
	}

	return "PoC_" + *name, *license, *remoteURL, *dir
}
