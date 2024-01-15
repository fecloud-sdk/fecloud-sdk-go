package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateBigkeyScanTaskRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o CreateBigkeyScanTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBigkeyScanTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateBigkeyScanTaskRequest", string(data)}, " ")
}
