// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kafka

import (
	"fmt"
	"path/filepath"
	"strings"
	"unicode"

	// Allow embedding bridge metadata
	_ "embed"

	"github.com/Mongey/terraform-provider-kafka/kafka"
	"github.com/pulumi/pulumi-kafka/provider/v3/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tks "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "kafka"

	// modules:
	mainMod = "index"
)

func ref[T any](t T) *T { return &t }

//go:embed cmd/pulumi-resource-kafka/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider.
func Provider() tfbridge.ProviderInfo {
	p := shimv2.NewProvider(kafka.Provider())
	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "kafka",
		GitHubOrg:         "Mongey",
		Description:       "A Pulumi package for creating and managing Kafka.",
		Keywords:          []string{"pulumi", "kafka"},
		License:           "Apache-2.0",
		Homepage:          "https://pulumi.io",
		Repository:        "https://github.com/pulumi/pulumi-kafka",
		Version:           version.Version,
		MetadataInfo:      tfbridge.NewProviderMetadata(metadata),
		TFProviderLicense: ref(tfbridge.MITLicenseType),
		Config: map[string]*tfbridge.SchemaInfo{
			"sasl_mechanism": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{
						"KAFKA_SASL_MECHANISM",
					},
					Value: "plain",
				},
			},
			"skip_tls_verify": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{
						"KAFKA_SKIP_VERIFY",
					},
					Value: false,
				},
			},
			"tls_enabled": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{
						"KAFKA_ENABLE_TLS",
					},
					Value: true,
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"kafka_acl": {
				Fields: map[string]*tfbridge.SchemaInfo{
					"resource_name": {
						Name: "aclResourceName",
					},
					"resource_type": {
						Name: "aclResourceType",
					},
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"kafka_topic": {Docs: noDocs},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				mainPkg: "Kafka",
			},
		},
	}

	prov.MustComputeTokens(tks.SingleModule("kafka", mainMod,
		// kafka doesn't apply tf's normal submodule naming for functions.
		//
		// Function token:
		//
		//	kafka:index:getTopic
		//
		// Resource token:
		//
		//	kafka:index/topic:Topic
		func(module, name string) (string, error) {
			if !strings.HasPrefix(name, "get") {
				module += "/" + string(unicode.ToLower(rune(name[0]))) + name[1:]
			}
			return fmt.Sprintf("kafka:%s:%s", module, name), nil
		}))

	prov.SetAutonaming(255, "-")

	prov.MustApplyAutoAliases()

	return prov
}

var noDocs = &tfbridge.DocInfo{
	Markdown: []byte{' '},
}
