package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchTagActionRequest struct {
	InstanceId string `json:"instance_id"`

	Body *BatchOperateInstanceTagRequestBody `json:"body,omitempty"`
}

func (o BatchTagActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTagActionRequest struct{}"
	}

	return strings.Join([]string{"BatchTagActionRequest", string(data)}, " ")
}
