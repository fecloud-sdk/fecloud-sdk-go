package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResizeInstanceVolumeOption struct {
	Size string `json:"size"`

	GroupId *string `json:"group_id,omitempty"`

	NodeIds *[]string `json:"node_ids,omitempty"`
}

func (o ResizeInstanceVolumeOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceVolumeOption struct{}"
	}

	return strings.Join([]string{"ResizeInstanceVolumeOption", string(data)}, " ")
}
