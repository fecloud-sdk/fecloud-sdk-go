package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type VulHostInfo struct {
	HostId *string `json:"host_id,omitempty"`

	SeverityLevel *string `json:"severity_level,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	AgentId *string `json:"agent_id,omitempty"`

	CveNum *int32 `json:"cve_num,omitempty"`

	CveIdList *[]string `json:"cve_id_list,omitempty"`

	Status *string `json:"status,omitempty"`

	RepairCmd *string `json:"repair_cmd,omitempty"`

	AppPath *string `json:"app_path,omitempty"`

	RegionName *string `json:"region_name,omitempty"`

	PublicIp *string `json:"public_ip,omitempty"`

	PrivateIp *string `json:"private_ip,omitempty"`

	GroupId *string `json:"group_id,omitempty"`

	GroupName *string `json:"group_name,omitempty"`

	OsType *string `json:"os_type,omitempty"`

	AssetValue *string `json:"asset_value,omitempty"`

	IsAffectBusiness *bool `json:"is_affect_business,omitempty"`

	FirstScanTime *int64 `json:"first_scan_time,omitempty"`

	ScanTime *int64 `json:"scan_time,omitempty"`
}

func (o VulHostInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulHostInfo struct{}"
	}

	return strings.Join([]string{"VulHostInfo", string(data)}, " ")
}
