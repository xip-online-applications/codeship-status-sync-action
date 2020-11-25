package watcher

import (
	"fmt"

	"github-codeship-connector/build"
)

type ResultHandler struct {
	build *build.Build
	err   error
}

func HandleBuildResult(build *build.Build, err error) {
	handler := &ResultHandler{
		build: build,
		err:   err,
	}

	handler.handle()
}

func (h *ResultHandler) handle() {
	if h.err == nil && h.build.HasSucceeded() {
		h.success()
		return
	}

	h.error()
}

func (h *ResultHandler) success() {
	fmt.Println(fmt.Sprintf("Succesfully handled the build: %s", h.build.Uuid()))
}

func (h *ResultHandler) error() {
	if h.err != nil {
		fmt.Println(fmt.Sprintf("Failed to handle the result and got an error: %s", h.err.Error()))
	} else {
		fmt.Println(fmt.Sprintf("Failed to handle the result with status: %s", h.build.Status()))
	}
}
