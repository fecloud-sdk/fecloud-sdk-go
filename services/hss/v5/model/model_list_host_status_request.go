package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListHostStatusRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Version *string `json:"version,omitempty"`

	AgentStatus *string `json:"agent_status,omitempty"`

	DetectResult *string `json:"detect_result,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	HostId *string `json:"host_id,omitempty"`

	HostStatus *string `json:"host_status,omitempty"`

	OsType *string `json:"os_type,omitempty"`

	PrivateIp *string `json:"private_ip,omitempty"`

	PublicIp *string `json:"public_ip,omitempty"`

	IpAddr *string `json:"ip_addr,omitempty"`

	ProtectStatus *string `json:"protect_status,omitempty"`

	GroupId *string `json:"group_id,omitempty"`

	GroupName *string `json:"group_name,omitempty"`

	Region *string `json:"region,omitempty"`

	PolicyGroupId *string `json:"policy_group_id,omitempty"`

	PolicyGroupName *string `json:"policy_group_name,omitempty"`

	ChargingMode *string `json:"charging_mode,omitempty"`

	Refresh *bool `json:"refresh,omitempty"`

	AboveVersion *bool `json:"above_version,omitempty"`

	OutsideHost *bool `json:"outside_host,omitempty"`

	AssetValue *string `json:"asset_value,omitempty"`

	Label *string `json:"label,omitempty"`

	ServerGroup *string `json:"server_group,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListHostStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostStatusRequest struct{}"
	}

	return strings.Join([]string{"ListHostStatusRequest", string(data)}, " ")
}
