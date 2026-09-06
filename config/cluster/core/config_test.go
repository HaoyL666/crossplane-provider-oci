/*
 * Copyright (c) 2026 Oracle and/or its affiliates
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package core

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
)

func TestConfigureComputeClusterOverridesPlacementConstraintNames(t *testing.T) {
	r := &config.Resource{References: config.References{}}
	configureComputeCluster(r)

	want := map[string]string{
		"PlacementConstraintDetailsInitParameters": "ComputeClusterPlacementConstraintDetailsInitParameters",
		"PlacementConstraintDetailsObservation":    "ComputeClusterPlacementConstraintDetailsObservation",
		"PlacementConstraintDetailsParameters":     "ComputeClusterPlacementConstraintDetailsParameters",
	}
	for old, expected := range want {
		if got := r.OverrideFieldNames[old]; got != expected {
			t.Errorf("OverrideFieldNames[%q] = %q, want %q", old, got, expected)
		}
	}
}
