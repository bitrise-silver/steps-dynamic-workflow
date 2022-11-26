package main

import (
	"bytes"
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/hairyhenderson/gomplate/v3"
)

func Generate(templateFile string, valueFile string) (*bytes.Buffer, error) {
	templateContent, err := os.ReadFile(templateFile)
	if err != nil {
		return nil, fmt.Errorf("error while reading template file: %s", err)
	}
	template := gomplate.Template{
		Name:   filename(templateFile),
		Text:   string(templateContent),
		Writer: &bytes.Buffer{},
	}
	options := gomplate.Options{
		Datasources: map[string]gomplate.Datasource{
			filename(valueFile): {URL: &url.URL{Scheme: "file", Path: valueFile}},
		},
	}

	ctx := context.Background()
	templateRender := gomplate.NewRenderer(options)
	err = templateRender.RenderTemplates(ctx, []gomplate.Template{template})
	if err != nil {
		return nil, fmt.Errorf("error while rendering template: %s", err)
	}
	return template.Writer.(*bytes.Buffer), nil
}
