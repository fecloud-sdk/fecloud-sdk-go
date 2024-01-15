package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type InstanceItem struct {
	InstanceId string `json:"instance_id"`

	InstanceName string `json:"instance_name"`

	Tags []InstanceItemTagItem `json:"tags"`
}

func (o InstanceItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceItem struct{}"
	}

	return strings.Join([]string{"InstanceItem", string(data)}, " ")
}
