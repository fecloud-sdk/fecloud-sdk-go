package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Host struct {
	HostName *string `json:"host_name,omitempty"`

	HostId *string `json:"host_id,omitempty"`

	AgentId *string `json:"agent_id,omitempty"`

	PrivateIp *string `json:"private_ip,omitempty"`

	PublicIp *string `json:"public_ip,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	HostStatus *string `json:"host_status,omitempty"`

	AgentStatus *string `json:"agent_status,omitempty"`

	InstallResultCode *string `json:"install_result_code,omitempty"`

	Version *string `json:"version,omitempty"`

	ProtectStatus *string `json:"protect_status,omitempty"`

	OsImage *string `json:"os_image,omitempty"`

	OsType *string `json:"os_type,omitempty"`

	OsBit *string `json:"os_bit,omitempty"`

	DetectResult *string `json:"detect_result,omitempty"`

	ChargingMode *string `json:"charging_mode,omitempty"`

	ResourceId *string `json:"resource_id,omitempty"`

	OutsideHost *bool `json:"outside_host,omitempty"`

	GroupId *string `json:"group_id,omitempty"`

	GroupName *string `json:"group_name,omitempty"`

	PolicyGroupId *string `json:"policy_group_id,omitempty"`

	PolicyGroupName *string `json:"policy_group_name,omitempty"`

	Asset *int32 `json:"asset,omitempty"`

	Vulnerability *int32 `json:"vulnerability,omitempty"`

	Baseline *int32 `json:"baseline,omitempty"`

	Intrusion *int32 `json:"intrusion,omitempty"`

	AssetValue *string `json:"asset_value,omitempty"`

	Labels *[]string `json:"labels,omitempty"`

	AgentCreateTime *int64 `json:"agent_create_time,omitempty"`

	AgentUpdateTime *int64 `json:"agent_update_time,omitempty"`

	AgentVersion *string `json:"agent_version,omitempty"`

	UpgradeStatus *string `json:"upgrade_status,omitempty"`

	UpgradeResultCode *string `json:"upgrade_result_code,omitempty"`

	Upgradable *bool `json:"upgradable,omitempty"`

	OpenTime *int64 `json:"open_time,omitempty"`
}

func (o Host) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Host struct{}"
	}

	return strings.Join([]string{"Host", string(data)}, " ")
}
