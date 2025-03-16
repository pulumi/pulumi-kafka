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
	"context"
	"fmt"
	"strings"

	original "github.com/Mongey/terraform-provider-kafka/kafka"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ForceDeleteProviderFactoryWrapper wraps the original provider factory and adds force_delete_resources option
func ForceDeleteProviderFactoryWrapper() *schema.Provider {
	originalProvider := original.Provider()
	
	// Add the force_delete_resources config option
	originalProvider.Schema["force_delete_resources"] = &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		DefaultFunc: schema.EnvDefaultFunc("KAFKA_FORCE_DELETE_RESOURCES", false),
		Description: "Force delete Kafka resources even if the Kafka cluster is not responsive",
	}

	// Save the original create/read/update/delete functions
	resources := map[string]*schema.Resource{}
	for resourceName, resource := range originalProvider.ResourcesMap {
		origDeleteContext := resource.DeleteContext
		
		// Wrap the delete operation with our force delete capability
		resources[resourceName] = &schema.Resource{
			Schema:       resource.Schema,
			CreateContext: resource.CreateContext,
			ReadContext:   resource.ReadContext,
			UpdateContext: resource.UpdateContext,
			DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
				// Get the force_delete_resources configuration
				forceDelete := false
				if config, ok := m.(*schema.ResourceData); ok {
					forceDelete = config.Get("force_delete_resources").(bool)
				}

				// Call the original delete function
				diags := origDeleteContext(ctx, d, m)
				
				// If there's an error and force_delete_resources is true, handle it
				if diags.HasError() && forceDelete {
					// Check if the error is related to Kafka connectivity
					errorMsg := diagsToString(diags)
					if isKafkaConnectivityError(errorMsg) {
						// Log the forced deletion
						fmt.Printf("[WARN] Force deleting resource %s due to Kafka connectivity issues. Error: %s\n", 
							d.Id(), errorMsg)
						
						// Clear the diagnostics and mark the resource as deleted
						d.SetId("")
						return nil
					}
				}
				
				return diags
			},
			Importer:     resource.Importer,
			Timeouts:     resource.Timeouts,
			CustomizeDiff: resource.CustomizeDiff,
			Description:  resource.Description,
			DeprecationMessage: resource.DeprecationMessage,
		}
	}

	// Replace the resources with our wrapped versions
	originalProvider.ResourcesMap = resources
	
	return originalProvider
}

// isKafkaConnectivityError checks if the error is related to Kafka connectivity
func isKafkaConnectivityError(errorMsg string) bool {
	errorPatterns := []string{
		"connection refused",
		"no such host",
		"i/o timeout",
		"cannot connect to kafka",
		"broker not available",
		"kafka server: Request timed out",
		"kafka cluster is not responsive",
		"failed to connect to kafka",
	}

	errorMsg = strings.ToLower(errorMsg)
	for _, pattern := range errorPatterns {
		if strings.Contains(errorMsg, pattern) {
			return true
		}
	}
	return false
}

// diagsToString converts diagnostics to a string
func diagsToString(diags diag.Diagnostics) string {
	var builder strings.Builder
	for _, diag := range diags {
		if diag.Severity == diag.Error {
			builder.WriteString(diag.Summary)
			builder.WriteString(": ")
			builder.WriteString(diag.Detail)
			builder.WriteString("\n")
		}
	}
	return builder.String()
} 