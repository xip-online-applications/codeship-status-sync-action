package watcher

import (
	"fmt"

	"github-codeship-connector/build"
	"github-codeship-connector/organization"
)

type Service struct {
	organizationService *organization.Service
}

func New(orgService *organization.Service) *Service {
	return &Service{
		organizationService: orgService,
	}
}

func (w *Service) Start() (*build.Build, error) {
	buildId := w.organizationService.BuildId()

	buildObj, err := build.FindAndWait(w.organizationService, buildId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Failed to find the build with ID '%s' we needed in Codeship due to: %s", buildId, err.Error()))
	}

	wrappedBuild, err := build.WaitUntilBuildInEndState(w.organizationService, buildObj)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Build with ID '%s' did not go to the proper state: %s", buildId, err.Error()))
	}

	return wrappedBuild, nil
}
