// Code generated by go generate; DO NOT EDIT.
// This file was generated at
// Wed Apr 29 16:31:41 UTC 2020
package main

import "github.com/hashicorp/nomad/nomad/structs"

var msgTypeNames = map[structs.MessageType]string{
	structs.NodeRegisterRequestType:                      "NodeRegisterRequestType",
	structs.NodeDeregisterRequestType:                    "NodeDeregisterRequestType",
	structs.NodeUpdateStatusRequestType:                  "NodeUpdateStatusRequestType",
	structs.NodeUpdateDrainRequestType:                   "NodeUpdateDrainRequestType",
	structs.JobRegisterRequestType:                       "JobRegisterRequestType",
	structs.JobDeregisterRequestType:                     "JobDeregisterRequestType",
	structs.EvalUpdateRequestType:                        "EvalUpdateRequestType",
	structs.EvalDeleteRequestType:                        "EvalDeleteRequestType",
	structs.AllocUpdateRequestType:                       "AllocUpdateRequestType",
	structs.AllocClientUpdateRequestType:                 "AllocClientUpdateRequestType",
	structs.ReconcileJobSummariesRequestType:             "ReconcileJobSummariesRequestType",
	structs.VaultAccessorRegisterRequestType:             "VaultAccessorRegisterRequestType",
	structs.VaultAccessorDeregisterRequestType:           "VaultAccessorDeregisterRequestType",
	structs.ApplyPlanResultsRequestType:                  "ApplyPlanResultsRequestType",
	structs.DeploymentStatusUpdateRequestType:            "DeploymentStatusUpdateRequestType",
	structs.DeploymentPromoteRequestType:                 "DeploymentPromoteRequestType",
	structs.DeploymentAllocHealthRequestType:             "DeploymentAllocHealthRequestType",
	structs.DeploymentDeleteRequestType:                  "DeploymentDeleteRequestType",
	structs.JobStabilityRequestType:                      "JobStabilityRequestType",
	structs.ACLPolicyUpsertRequestType:                   "ACLPolicyUpsertRequestType",
	structs.ACLPolicyDeleteRequestType:                   "ACLPolicyDeleteRequestType",
	structs.ACLTokenUpsertRequestType:                    "ACLTokenUpsertRequestType",
	structs.ACLTokenDeleteRequestType:                    "ACLTokenDeleteRequestType",
	structs.ACLTokenBootstrapRequestType:                 "ACLTokenBootstrapRequestType",
	structs.AutopilotRequestType:                         "AutopilotRequestType",
	structs.UpsertNodeEventsType:                         "UpsertNodeEventsType",
	structs.JobBatchDeregisterRequestType:                "JobBatchDeregisterRequestType",
	structs.AllocUpdateDesiredTransitionRequestType:      "AllocUpdateDesiredTransitionRequestType",
	structs.NodeUpdateEligibilityRequestType:             "NodeUpdateEligibilityRequestType",
	structs.BatchNodeUpdateDrainRequestType:              "BatchNodeUpdateDrainRequestType",
	structs.SchedulerConfigRequestType:                   "SchedulerConfigRequestType",
	structs.NodeBatchDeregisterRequestType:               "NodeBatchDeregisterRequestType",
	structs.ClusterMetadataRequestType:                   "ClusterMetadataRequestType",
	structs.ServiceIdentityAccessorRegisterRequestType:   "ServiceIdentityAccessorRegisterRequestType",
	structs.ServiceIdentityAccessorDeregisterRequestType: "ServiceIdentityAccessorDeregisterRequestType",
	structs.CSIVolumeRegisterRequestType:                 "CSIVolumeRegisterRequestType",
	structs.CSIVolumeDeregisterRequestType:               "CSIVolumeDeregisterRequestType",
	structs.CSIVolumeClaimRequestType:                    "CSIVolumeClaimRequestType",
	structs.ScalingEventRegisterRequestType:              "ScalingEventRegisterRequestType",
}
