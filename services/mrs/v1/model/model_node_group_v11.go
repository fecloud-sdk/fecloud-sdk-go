package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NodeGroupV11 struct {
	GroupName string `json:"group_name"`

	NodeNum int32 `json:"node_num"`

	NodeSize string `json:"node_size"`

	RootVolumeSize *string `json:"root_volume_size,omitempty"`

	RootVolumeType *string `json:"root_volume_type,omitempty"`

	DataVolumeType *string `json:"data_volume_type,omitempty"`

	DataVolumeCount *int32 `json:"data_volume_count,omitempty"`

	DataVolumeSize *int32 `json:"data_volume_size,omitempty"`

	AutoScalingPolicy *AutoScalingPolicy `json:"auto_scaling_policy,omitempty"`
}

func (o NodeGroupV11) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeGroupV11 struct{}"
	}

	return strings.Join([]string{"NodeGroupV11", string(data)}, " ")
}
