package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NodeGroupV2 struct {
	GroupName string `json:"group_name"`

	NodeNum int32 `json:"node_num"`

	NodeSize string `json:"node_size"`

	RootVolume *Volume `json:"root_volume,omitempty"`

	DataVolume *Volume `json:"data_volume,omitempty"`

	DataVolumeCount *int32 `json:"data_volume_count,omitempty"`

	AutoScalingPolicy *AutoScalingPolicy `json:"auto_scaling_policy,omitempty"`

	AssignedRoles *[]string `json:"assigned_roles,omitempty"`
}

func (o NodeGroupV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeGroupV2 struct{}"
	}

	return strings.Join([]string{"NodeGroupV2", string(data)}, " ")
}
