package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateHotkeyScanTaskRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o CreateHotkeyScanTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHotkeyScanTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateHotkeyScanTaskRequest", string(data)}, " ")
}
