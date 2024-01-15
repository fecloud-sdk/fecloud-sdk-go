package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteGroupRequest struct {
	InstanceId string `json:"instance_id"`

	Body *BatchDeleteGroupReq `json:"body,omitempty"`
}

func (o BatchDeleteGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteGroupRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteGroupRequest", string(data)}, " ")
}
