package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateAutoExpireScanTaskRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o CreateAutoExpireScanTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAutoExpireScanTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateAutoExpireScanTaskRequest", string(data)}, " ")
}
