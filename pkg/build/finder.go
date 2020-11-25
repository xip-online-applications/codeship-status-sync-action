package build

import (
	"context"
	"fmt"
	"time"

	"github.com/codeship/codeship-go"

	"github-codeship-connector/organization"
)

type Finder struct {
	organizationService *organization.Service
	build               *codeship.Build
}

func FindAndWait(organizationService *organization.Service, buildId string) (*Build, error) {
	service := &Finder{
		organizationService: organizationService,
	}

	codeshipBuild, err := service.findAndWait(buildId)
	if err != nil {
		return nil, err
	}

	return fromCodeshipBuild(codeshipBuild), nil
}

func (s *Finder) findAndWait(buildId string) (*codeship.Build, error) {
	build := (*codeship.Build)(nil)
	runCount := 0

	fmt.Println(fmt.Sprintf("Searching for the build with ID %s...", buildId))

	for build == nil {
		var err error = nil

		build, err = s.runFindIteration(buildId)
		if err != nil {
			return nil, err
		}
		if build != nil {
			continue
		}

		runCount++
		if runCount > 300 {
			return nil, fmt.Errorf("tried for at least 5 minutes, but failed")
		}

		time.Sleep(1 * time.Second)
	}

	return build, nil
}

func (s *Finder) runFindIteration(buildId string) (*codeship.Build, error) {
	currentBuildList, err := s.retrieveBuilds()
	if err != nil {
		return nil, nil
	}

	if currentBuildList.Total == 0 {
		return nil, fmt.Errorf("this project has no builds")
	}

	currentBuild := s.findBuildFromList(currentBuildList, buildId)
	if currentBuild == nil {
		return nil, nil
	}

	return currentBuild, nil
}

func (s *Finder) retrieveBuilds() (*codeship.BuildList, error) {
	projectUuid := s.organizationService.ProjectUuid()
	org := s.organizationService.Organization()

	builds, _, err := org.ListBuilds(context.Background(), projectUuid)
	if err != nil {
		return nil, err
	}

	return &builds, nil
}

func (s *Finder) findBuildFromList(builds *codeship.BuildList, buildId string) *codeship.Build {
	for _, build := range builds.Builds {
		if build.CommitSha == buildId {
			return &build
		}
	}

	return nil
}
