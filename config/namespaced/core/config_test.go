package core

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
)

func TestConfigureComputeClusterPlacementConstraintTypeNames(t *testing.T) {
	r := &config.Resource{OverrideFieldNames: map[string]string{}}

	configureComputeClusterPlacementConstraintTypeNames(r)

	want := map[string]string{
		"PlacementConstraintDetailsInitParameters": "ComputeClusterPlacementConstraintDetailsInitParameters",
		"PlacementConstraintDetailsParameters":     "ComputeClusterPlacementConstraintDetailsParameters",
		"PlacementConstraintDetailsObservation":    "ComputeClusterPlacementConstraintDetailsObservation",
	}
	for generatedName, wantName := range want {
		if got := r.OverrideFieldNames[generatedName]; got != wantName {
			t.Errorf("OverrideFieldNames[%q] = %q, want %q", generatedName, got, wantName)
		}
	}
}
