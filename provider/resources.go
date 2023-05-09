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
	// embed is used to store bridge-metadata.json in the compiled binary
	_ "embed"
	"path/filepath"
	"unicode"

	"github.com/Mongey/terraform-provider-kafka/kafka"
	"github.com/pulumi/pulumi-kafka/provider/v3/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/x"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "kafka"

	// modules:
	mainMod = "index"
)

func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

func makeDataSource(mod string, res string) tokens.ModuleMember {
	return makeMember(mod, res)
}

func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

func tfLicenseTypeRef(license tfbridge.TFProviderLicense) *tfbridge.TFProviderLicense {
	return &license
}

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
		TFProviderLicense: tfLicenseTypeRef(tfbridge.MITLicenseType),
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
				Tok: makeResource(mainMod, "Acl"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"resource_name": {
						Name: "aclResourceName",
					},
					"resource_type": {
						Name: "aclResourceType",
					},
				},
			},
			"kafka_topic": {Tok: makeResource(mainMod, "Topic")},
			"kafka_quota": {Tok: makeResource(mainMod, "Quota")},
			"kafka_user_scram_credential": {
				Tok:  makeResource(mainMod, "UserScramCredential"),
				Docs: &tfbridge.DocInfo{Markdown: []byte{' '}},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"kafka_topic": {
				Tok: makeDataSource(mainMod, "getTopic"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
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
		}, MetadataInfo: tfbridge.NewProviderMetadata(metadata),
	}

	err := x.ComputeDefaults(&prov, x.TokensSingleModule("kafka_", mainMod,
		x.MakeStandardToken(mainPkg)))
	contract.AssertNoErrorf(err, "failed to compute defaults")
	err = x.AutoAliasing(&prov, prov.GetMetadata())
	contract.AssertNoErrorf(err, "auto aliasing apply failed")

	prov.SetAutonaming(255, "-")

	return prov
}

//go:embed cmd/pulumi-resource-kafka/bridge-metadata.json
var metadata []byte
