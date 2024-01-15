package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type InstanceResult struct {
	InstanceId string `json:"instance_id"`

	InstanceName string `json:"instance_name"`

	Tags []InstanceTagResult `json:"tags"`
}

func (o InstanceResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceResult struct{}"
	}

	return strings.Join([]string{"InstanceResult", string(data)}, " ")
}
