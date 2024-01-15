package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchRestartOrDeleteInstancesRequest struct {
	Body *BatchRestartOrDeleteInstanceReq `json:"body,omitempty"`
}

func (o BatchRestartOrDeleteInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRestartOrDeleteInstancesRequest struct{}"
	}

	return strings.Join([]string{"BatchRestartOrDeleteInstancesRequest", string(data)}, " ")
}
