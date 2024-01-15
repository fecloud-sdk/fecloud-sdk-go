package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ProtectionServerInfo struct {
	HostId *string `json:"host_id,omitempty"`

	AgentId *string `json:"agent_id,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	PrivateIp *string `json:"private_ip,omitempty"`

	OsType *string `json:"os_type,omitempty"`

	OsName *string `json:"os_name,omitempty"`

	HostStatus *string `json:"host_status,omitempty"`

	RansomProtectionStatus *string `json:"ransom_protection_status,omitempty"`

	ProtectStatus *string `json:"protect_status,omitempty"`

	GroupId *string `json:"group_id,omitempty"`

	GroupName *string `json:"group_name,omitempty"`

	ProtectPolicyId *string `json:"protect_policy_id,omitempty"`

	ProtectPolicyName *string `json:"protect_policy_name,omitempty"`

	BackupError *ProtectionServerInfoBackupError `json:"backup_error,omitempty"`

	BackupProtectionStatus *string `json:"backup_protection_status,omitempty"`

	CountProtectEvent *int32 `json:"count_protect_event,omitempty"`

	CountBackuped *int32 `json:"count_backuped,omitempty"`

	AgentStatus *string `json:"agent_status,omitempty"`

	Version *string `json:"version,omitempty"`

	VaultId *string `json:"vault_id,omitempty"`

	VaultName *string `json:"vault_name,omitempty"`

	VaultSize *int32 `json:"vault_size,omitempty"`

	VaultUsed *int32 `json:"vault_used,omitempty"`

	VaultAllocated *int32 `json:"vault_allocated,omitempty"`

	VaultChargingMode *string `json:"vault_charging_mode,omitempty"`

	VaultStatus *string `json:"vault_status,omitempty"`

	BackupPolicyId *string `json:"backup_policy_id,omitempty"`

	BackupPolicyName *string `json:"backup_policy_name,omitempty"`

	BackupPolicyEnabled *bool `json:"backup_policy_enabled,omitempty"`

	ResourcesNum *int32 `json:"resources_num,omitempty"`
}

func (o ProtectionServerInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectionServerInfo struct{}"
	}

	return strings.Join([]string{"ProtectionServerInfo", string(data)}, " ")
}
