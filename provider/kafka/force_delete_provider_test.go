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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/stretchr/testify/assert"
)

func TestIsKafkaConnectivityError(t *testing.T) {
	testCases := []struct {
		name     string
		errorMsg string
		expected bool
	}{
		{
			name:     "connection refused error",
			errorMsg: "Error: connection refused",
			expected: true,
		},
		{
			name:     "no such host error",
			errorMsg: "Error: no such host",
			expected: true,
		},
		{
			name:     "i/o timeout error",
			errorMsg: "Error: i/o timeout",
			expected: true,
		},
		{
			name:     "cannot connect to kafka error",
			errorMsg: "Error: cannot connect to kafka",
			expected: true,
		},
		{
			name:     "broker not available error",
			errorMsg: "Error: broker not available",
			expected: true,
		},
		{
			name:     "kafka server timeout error",
			errorMsg: "Error: kafka server: Request timed out",
			expected: true,
		},
		{
			name:     "kafka cluster not responsive error",
			errorMsg: "Error: kafka cluster is not responsive",
			expected: true,
		},
		{
			name:     "failed to connect to kafka error",
			errorMsg: "Error: failed to connect to kafka",
			expected: true,
		},
		{
			name:     "unrelated error",
			errorMsg: "Error: invalid configuration",
			expected: false,
		},
		{
			name:     "empty error",
			errorMsg: "",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isKafkaConnectivityError(tc.errorMsg)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestDiagsToString(t *testing.T) {
	diags := diag.Diagnostics{
		diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Summary",
			Detail:   "Error Detail",
		},
		diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Warning Summary",
			Detail:   "Warning Detail",
		},
	}

	result := diagsToString(diags)
	assert.Contains(t, result, "Error Summary")
	assert.Contains(t, result, "Error Detail")
	assert.NotContains(t, result, "Warning Summary")
	assert.NotContains(t, result, "Warning Detail")
} 