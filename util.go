package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/bitrise-io/go-utils/log"
)

func filename(path string) string {
	return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
}

func executeCommand(command ...string) string {
	output, err := exec.Command(command[0], command[1:]...).CombinedOutput()
	if err != nil {
		failf("Error running command %s . Output: %s", command, err)
	}
	return string(output)
}

func executeCommandInDir(cwd string, command ...string) string {
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Dir = cwd
	output, err := cmd.CombinedOutput()
	if err != nil {
		failf("Error running command %s . Output: %s", command, err)
	}
	return string(output)
}

func failf(format string, v ...interface{}) {
	log.Errorf(format, v...)
	os.Exit(1)
}
