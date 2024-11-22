package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func createGitHubRepo(repoName, orgName, token string) error {
	client := resty.New()

	// GitHub APIリクエスト
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+token).
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetBody(map[string]interface{}{
			"name": repoName,
		}).
		Post(fmt.Sprintf("https://api.github.com/orgs/%s/repos", orgName))

	if err != nil || resp.StatusCode() != 201 {
		return fmt.Errorf("GitHub API request failed: %v, response: %s", err, resp.String())
	}
	return nil
}
