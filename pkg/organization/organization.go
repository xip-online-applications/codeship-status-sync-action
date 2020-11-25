package organization

import (
	"context"

	"github.com/codeship/codeship-go"
)

type Service struct {
	config *Config

	client       *codeship.Client
	organization *codeship.Organization
}

func NewFromEnvironmentVariables() *Service {
	config := ConfigForGithubAction()

	return New(config)
}

func New(config *Config) *Service {
	service := &Service{
		config: config,
	}

	service.authenticate(config.username, config.password)
	service.retrieveOrganization()

	return service
}

func (s *Service) Organization() *codeship.Organization {
	return s.organization
}

func (s *Service) ProjectUuid() string {
	return s.config.projectUuid
}

func (s *Service) BuildId() string {
	return s.config.gitCommitSha
}

func (s *Service) authenticate(username string, password string) {
	auth := codeship.NewBasicAuth(username, password)
	client, err := codeship.New(auth)
	if err != nil {
		panic("Failed to authenticate to codeship")
	}

	s.client = client
}

func (s *Service) retrieveOrganization() {
	org, err := s.client.Organization(context.Background(), s.config.organizationUuid)
	if err != nil {
		panic("Request organization not found")
	}

	s.organization = org
}
