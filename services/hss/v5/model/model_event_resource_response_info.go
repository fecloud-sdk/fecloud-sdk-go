package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EventResourceResponseInfo struct {
	DomainId *string `json:"domain_id,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	RegionName *string `json:"region_name,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	CloudId *string `json:"cloud_id,omitempty"`

	VmName *string `json:"vm_name,omitempty"`

	VmUuid *string `json:"vm_uuid,omitempty"`

	ContainerId *string `json:"container_id,omitempty"`

	ImageId *string `json:"image_id,omitempty"`

	ImageName *string `json:"image_name,omitempty"`

	HostAttr *string `json:"host_attr,omitempty"`

	Service *string `json:"service,omitempty"`

	MicroService *string `json:"micro_service,omitempty"`

	SysArch *string `json:"sys_arch,omitempty"`

	OsBit *string `json:"os_bit,omitempty"`

	OsType *string `json:"os_type,omitempty"`

	OsName *string `json:"os_name,omitempty"`

	OsVersion *string `json:"os_version,omitempty"`
}

func (o EventResourceResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventResourceResponseInfo struct{}"
	}

	return strings.Join([]string{"EventResourceResponseInfo", string(data)}, " ")
}
