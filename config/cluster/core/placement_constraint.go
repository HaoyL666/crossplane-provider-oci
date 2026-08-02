package core

import "github.com/crossplane/upjet/v2/pkg/config"

func configureComputeClusterPlacementConstraintTypeNames(r *config.Resource) {
	r.OverrideFieldNames["PlacementConstraintDetailsInitParameters"] = "ComputeClusterPlacementConstraintDetailsInitParameters"
	r.OverrideFieldNames["PlacementConstraintDetailsParameters"] = "ComputeClusterPlacementConstraintDetailsParameters"
	r.OverrideFieldNames["PlacementConstraintDetailsObservation"] = "ComputeClusterPlacementConstraintDetailsObservation"
}
