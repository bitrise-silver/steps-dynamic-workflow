package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/bitrise-io/go-utils/log"
	"github.com/google/uuid"
)

func triggerDynamicPipeline(config *bytes.Buffer, repo string, branch string, pipelineID string, appSlug string, token string, cleanUp bool) {
	// Git clone, commit and push
	tempFolder := uuid.NewString()
	executeCommand("git", "clone", repo, tempFolder)
	executeCommandInDir(tempFolder, "git", "checkout", "-b", branch)
	err := os.WriteFile(tempFolder+"/bitrise.yml", config.Bytes(), 0644)
	if err != nil {
		failf("Error while writing bitrise.yml: %s", err)
	}
	executeCommandInDir(tempFolder, "git", "add", "bitrise.yml")
	executeCommandInDir(tempFolder, "git", "commit", "-m", "Bitrise generated workflow")
	executeCommandInDir(tempFolder, "git", "push", "origin", branch)

	// Trigger
	params := fmt.Sprintf("{\"hook_info\":{\"type\":\"bitrise\"},\"build_params\":{\"branch\":\"%s\",\"pipeline_id\":\"%s\"}}", branch, pipelineID)
	triggerResponse := executeCommand("curl", "-sS", "-X", "POST", "-H", "Authorization: "+token, "https://api.bitrise.io/v0.1/apps/"+appSlug+"/builds", "-d", params)

	type BuildResponse struct {
		Status    string `json:"status,omitempty"`
		Message   string `json:"message,omitempty"`
		Slug      string `json:"slug,omitempty"`
		Service   string `json:"service,omitempty"`
		BuildSlug string `json:"build_slug,omitempty"`
		BuildURL  string `json:"build_url,omitempty"`

		BuildNumber       json.Number `json:"build_number,omitempty"`
		TriggeredWorkflow string      `json:"triggered_workflow,omitempty"`
	}

	var r BuildResponse
	if err := json.Unmarshal([]byte(triggerResponse), &r); err != nil {
		failf("Error parsing /builds response: %s", err)
	}
	log.Infof("Pipeline started: " + r.BuildURL)

	executeCommand("bitrise", "envman", "add", "--key", "DYNAMIC_TRIGGERED_BUILD_ID", "--value", r.BuildSlug)

	if cleanUp {
		log.Infof("Waiting for pipeline to finish...")
		app := NewAppWithDefaultURL(appSlug, token)
		err := app.WaitForPipeline(r.BuildSlug)
		executeCommandInDir(tempFolder, "git", "push", "origin", ":"+branch)
		executeCommand("rm", "-rf", tempFolder)
		if err != nil {
			failf("An error occurred: %s", err)
		}
	}
}
