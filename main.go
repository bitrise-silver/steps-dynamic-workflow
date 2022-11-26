package main

import (
	"os"

	"github.com/bitrise-io/go-steputils/stepconf"
)

type Config struct {
	TemplateFile     string          `env:"template_file,required"`
	ValueFile        string          `env:"value_file,required"`
	RepositoryURL    string          `env:"repository_url,required"`
	RepositoryBranch string          `env:"repository_branch,required"`
	CleanUpBranch    bool            `env:"clean_up_branch_after_used,required"`
	AccessToken      stepconf.Secret `env:"access_token,required"`
	AppSlug          string          `env:"BITRISE_APP_SLUG,required"`
	PipelineID       string          `env:"pipeline_id,required"`
}

func main() {
	var cfg Config
	if err := stepconf.Parse(&cfg); err != nil {
		failf("Error while reading template file: %s", err)
	}
	template, err := Generate(cfg.TemplateFile, cfg.ValueFile)
	if err != nil {
		failf("Error while generating template: %s", err)
	}
	triggerDynamicPipeline(template, cfg.RepositoryURL, cfg.RepositoryBranch, cfg.PipelineID, cfg.AppSlug, string(cfg.AccessToken), cfg.CleanUpBranch)
	os.Exit(0)
}
