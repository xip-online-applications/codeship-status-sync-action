package build

import (
	"context"
	"fmt"
	"time"

	"github-codeship-connector/organization"
)

type Waiter struct {
	organizationService *organization.Service
	build               *Build
}

func WaitUntilBuildInEndState(organizationService *organization.Service, build *Build) (*Build, error) {
	waiter := &Waiter{
		organizationService: organizationService,
		build:               build,
	}

	return waiter.waitUntilInEndState()
}

func (w *Waiter) waitUntilInEndState() (*Build, error) {
	runCount := 0

	fmt.Println(fmt.Sprintf("Waiting build with ID %s to be done...", w.build.Uuid()))

	for true {
		if !w.build.IsRunning() {
			return w.build, nil
		}

		runCount++
		if runCount > 300 {
			break
		}

		time.Sleep(1 * time.Second)
		w.retrieveAndUpdateInternalBuild()
	}

	return nil, fmt.Errorf("tried for at least 5 minutes, but failed")
}

func (w *Waiter) retrieveAndUpdateInternalBuild() {
	latestBuild, _ := w.buildGetLatest()
	if latestBuild == nil {
		return
	}

	w.build = latestBuild
}

func (w *Waiter) buildGetLatest() (*Build, error) {
	buildUuid := w.build.Uuid()
	projectUuid := w.organizationService.ProjectUuid()
	org := w.organizationService.Organization()

	latestBuild, _, err := org.GetBuild(context.Background(), projectUuid, buildUuid)
	if err != nil {
		return nil, err
	}

	return fromCodeshipBuild(&latestBuild), nil
}
