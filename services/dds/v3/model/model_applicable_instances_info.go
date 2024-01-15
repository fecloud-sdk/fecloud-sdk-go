package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ApplicableInstancesInfo struct {
	InstanceId string `json:"instance_id"`

	InstanceName string `json:"instance_name"`

	Entities []EntityInfo `json:"entities"`
}

func (o ApplicableInstancesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicableInstancesInfo struct{}"
	}

	return strings.Join([]string{"ApplicableInstancesInfo", string(data)}, " ")
}
