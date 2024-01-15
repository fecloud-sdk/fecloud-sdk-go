package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateProtectionPolicyInfoRequestInfo struct {
	PolicyId string `json:"policy_id"`

	PolicyName string `json:"policy_name"`

	ProtectionMode string `json:"protection_mode"`

	BaitProtectionStatus string `json:"bait_protection_status"`

	ProtectionDirectory string `json:"protection_directory"`

	ProtectionType string `json:"protection_type"`

	ExcludeDirectory *string `json:"exclude_directory,omitempty"`

	AgentIdList *[]string `json:"agent_id_list,omitempty"`

	OperatingSystem string `json:"operating_system"`

	RuntimeDetectionStatus *string `json:"runtime_detection_status,omitempty"`

	ProcessWhitelist *[]TrustProcessInfo `json:"process_whitelist,omitempty"`
}

func (o UpdateProtectionPolicyInfoRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProtectionPolicyInfoRequestInfo struct{}"
	}

	return strings.Join([]string{"UpdateProtectionPolicyInfoRequestInfo", string(data)}, " ")
}
