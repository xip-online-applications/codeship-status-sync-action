package main

import (
	"github-codeship-connector/organization"
	"github-codeship-connector/watcher"
)

func main() {
	orgService := organization.NewFromEnvironmentVariables()

	watcherService := watcher.New(orgService)
	result, err := watcherService.Start()

	watcher.HandleBuildResult(result, err)
}
