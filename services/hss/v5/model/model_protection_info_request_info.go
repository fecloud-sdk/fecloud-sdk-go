package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ProtectionInfoRequestInfo struct {
	OperatingSystem string `json:"operating_system"`

	RansomProtectionStatus string `json:"ransom_protection_status"`

	ProtectionPolicyId *string `json:"protection_policy_id,omitempty"`

	CreateProtectionPolicy *ProtectionProxyInfoRequestInfo `json:"create_protection_policy,omitempty"`

	BackupProtectionStatus string `json:"backup_protection_status"`

	BackupResources *BackupResources `json:"backup_resources,omitempty"`

	BackupPolicyId *string `json:"backup_policy_id,omitempty"`

	BackupCycle *UpdateBackupPolicyRequestInfo1 `json:"backup_cycle,omitempty"`

	AgentIdList []string `json:"agent_id_list"`

	HostIdList []string `json:"host_id_list"`
}

func (o ProtectionInfoRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectionInfoRequestInfo struct{}"
	}

	return strings.Join([]string{"ProtectionInfoRequestInfo", string(data)}, " ")
}
