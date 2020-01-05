// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/testing/integration"
)

func TestAccTopic(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "topic"),
		})

	integration.ProgramTest(t, &test)
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.Error("expected a valid working directory", err)
	}

	return cwd
}

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{}
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		Config: map[string]string{
			"kafka:bootstrapServers": "[\"localhost:9092\"]",
			"kafka:tlsEnabled": "false",
		},
		Dependencies: []string{
			"@pulumi/kafka",
		},
	})

	return baseJS
}
