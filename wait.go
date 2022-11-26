package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bitrise-io/go-utils/log"
)

type App struct {
	BaseURL, Slug, AccessToken string
	IsDebugRetryTimings        bool
}

type Pipeline struct {
	Slug   string `json:"slug"`
	Status string `json:"status"`
}

// NewAppWithDefaultURL returns a Bitrise client with the default URl
func NewAppWithDefaultURL(slug, accessToken string) App {
	return App{
		BaseURL:     "https://api.bitrise.io",
		Slug:        slug,
		AccessToken: accessToken,
	}
}

func (app App) GetPipeline(pipelineId string) (pipeline Pipeline, err error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v0.1/apps/%s/pipelines/%s", app.BaseURL, app.Slug, pipelineId), nil)
	if err != nil {
		return Pipeline{}, err
	}

	req.Header.Add("Authorization", "token "+app.AccessToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Pipeline{}, err
	}

	defer func() {
		if cerr := resp.Body.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Pipeline{}, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return Pipeline{}, fmt.Errorf("failed to get response, statuscode: %d, body: %s", resp.StatusCode, respBody)
	}

	var response Pipeline
	if err := json.Unmarshal(respBody, &response); err != nil {
		return Pipeline{}, fmt.Errorf("failed to decode response, body: %s, error: %s", respBody, err)
	}
	return response, nil
}

func (app App) WaitForPipeline(slug string) error {
	for {
		pipeline, err := app.GetPipeline(slug)
		if err != nil {
			return fmt.Errorf("failed to get build info, error: %s", err)
		}
		if pipeline.Status == "initializing" || pipeline.Status == "running" || pipeline.Status == "on_hold" {
			time.Sleep(time.Second * 3)
			continue
		}
		log.Infof("pipeline finished with status: %s", pipeline.Status)
		break
	}
	return nil
}
