package organization

import (
	"os"
)

type Config struct {
	username         string
	password         string
	organizationUuid string
	projectUuid      string
	gitCommitSha     string
}

func ConfigForGithubAction() *Config {
	username := os.Getenv("INPUT_CODESHIPUSERNAME")
	password := os.Getenv("INPUT_CODESHIPPASSWORD")
	organizationUuid := os.Getenv("INPUT_CODESHIPORGANIZATION")
	projectUuid := os.Getenv("INPUT_CODESHIPPROJECTUUID")
	gitCommitSha := os.Getenv("INPUT_GITCOMMITSHA")

	return &Config{
		username:         username,
		password:         password,
		organizationUuid: organizationUuid,
		projectUuid:      projectUuid,
		gitCommitSha:     gitCommitSha,
	}
}
