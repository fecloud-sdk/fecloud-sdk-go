package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ProtectionProxyInfoRequestInfo struct {
	PolicyId *string `json:"policy_id,omitempty"`

	PolicyName *string `json:"policy_name,omitempty"`

	ProtectionMode *string `json:"protection_mode,omitempty"`

	BaitProtectionStatus *string `json:"bait_protection_status,omitempty"`

	ProtectionDirectory *string `json:"protection_directory,omitempty"`

	ProtectionType *string `json:"protection_type,omitempty"`

	ExcludeDirectory *string `json:"exclude_directory,omitempty"`

	RuntimeDetectionStatus *string `json:"runtime_detection_status,omitempty"`

	OperatingSystem *string `json:"operating_system,omitempty"`

	ProcessWhitelist *[]TrustProcessInfo `json:"process_whitelist,omitempty"`
}

func (o ProtectionProxyInfoRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectionProxyInfoRequestInfo struct{}"
	}

	return strings.Join([]string{"ProtectionProxyInfoRequestInfo", string(data)}, " ")
}
