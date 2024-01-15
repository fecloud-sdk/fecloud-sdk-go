package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type TargetInstanceRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o TargetInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetInstanceRequest struct{}"
	}

	return strings.Join([]string{"TargetInstanceRequest", string(data)}, " ")
}
