package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type HostModel struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Ip *string `json:"ip,omitempty"`

	AvailabilityZoneId *string `json:"availability_zone_id,omitempty"`

	Tags *[]TagPlain `json:"tags,omitempty"`

	Status *string `json:"status,omitempty"`

	ResourceId *string `json:"resource_id,omitempty"`

	Flavor *string `json:"flavor,omitempty"`

	Type *string `json:"type,omitempty"`

	Mem *string `json:"mem,omitempty"`

	Cpu *string `json:"cpu,omitempty"`

	RootVolumeSize *string `json:"root_volume_size,omitempty"`

	DataVolumeType *string `json:"data_volume_type,omitempty"`

	DataVolumeSize *int32 `json:"data_volume_size,omitempty"`

	DataVolumeCount *int32 `json:"data_volume_count,omitempty"`
}

func (o HostModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostModel struct{}"
	}

	return strings.Join([]string{"HostModel", string(data)}, " ")
}
