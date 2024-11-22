package main

import (
	"log"
	"os"
)

func main() {
	// 引数処理
	repoName, license, remoteURL, targetDir := parseCLIArgs()

	// 環境変数から組織名とトークンを取得
	orgName := os.Getenv("GITHUB_ORG")
	token := os.Getenv("GITHUB_TOKEN")
	if orgName == "" || token == "" {
		log.Fatal("Error: GITHUB_ORG and GITHUB_TOKEN environment variables must be set")
	}

	// GitHubリポジトリ作成
	err := createGitHubRepo(repoName, orgName, token)
	if err != nil {
		log.Fatalf("Failed to create GitHub repo: %v\n", err)
	}

	// ローカルリポジトリ作成
	err = initLocalRepo(repoName, license, orgName, remoteURL, targetDir)
	if err != nil {
		log.Fatalf("Failed to initialize local repo: %v\n", err)
	}

	log.Println("PoC repository setup completed successfully.")
}
