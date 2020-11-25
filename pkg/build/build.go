package build

import (
	"github.com/codeship/codeship-go"
)

type Build struct {
	build *codeship.Build
}

func fromCodeshipBuild(build *codeship.Build) *Build {
	return &Build{
		build: build,
	}
}

func (b *Build) Uuid() string {
	return b.build.UUID
}

func (b *Build) Status() string {
	return b.build.Status
}

func (b *Build) IsRunning() bool {
	state := b.build.Status

	switch state {
	case
		"waiting",
		"testing":
		return true
	}

	return false
}

func (b *Build) HasFailed() bool {
	state := b.build.Status

	switch state {
	case
		"failed",
		"paused":
		return true
	}

	return false
}

func (b *Build) HasBeenStopped() bool {
	state := b.build.Status

	switch state {
	case
		"stopping",
		"stopped":
		return true
	}

	return false
}

func (b *Build) HasSucceeded() bool {
	state := b.build.Status

	switch state {
	case
		"success":
		return true
	}

	return false
}
