package main

import (
	"embed"
	"fmt"
)

// ライセンス埋め込み
//go:embed licenses/*
var licenseFiles embed.FS

func getLicenseContent(license string) (string, error) {
	var fileName string
	switch license {
	case "MIT":
		fileName = "licenses/MIT.txt"
	case "Unlicense":
		fileName = "licenses/Unlicense.txt"
	default:
		return "", fmt.Errorf("unknown license type: %s", license)
	}

	content, err := licenseFiles.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to read license file: %w", err)
	}
	return string(content), nil
}
