package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteInstancesRequest struct {
	Body *BatchDeleteInstanceReq `json:"body,omitempty"`
}

func (o BatchDeleteInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstancesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstancesRequest", string(data)}, " ")
}
