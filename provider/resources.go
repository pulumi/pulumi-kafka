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
	"unicode"

	"github.com/Mongey/terraform-provider-kafka/kafka"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfshim"
	shimv1 "github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfshim/sdk-v1"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
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

func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

func tfLicenseTypeRef(license tfbridge.TFProviderLicense) *tfbridge.TFProviderLicense {
	return &license
}

// Provider returns additional overlaid schema and metadata associated with the provider.
func Provider() tfbridge.ProviderInfo {
	p := shimv1.NewProvider(kafka.Provider().(*schema.Provider))
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
			"ca_cert": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{
						"KAFKA_CA_CERT",
					},
				},
			},
			"client_cert": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{
						"KAFKA_CLIENT_CERT",
					},
				},
			},
			"client_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{
						"KAFKA_CLIENT_KEY",
					},
				},
			},
			"sasl_username": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{
						"KAFKA_SASL_USERNAME",
					},
				},
			},
			"sasl_password": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{
						"KAFKA_SASL_PASSWORD",
					},
				},
			},
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
		PreConfigureCallback: preConfigureCallback,
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
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^2.15.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=2.15.0,<3.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "2.*",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: map[string]string{
				mainPkg: "Kafka",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
