package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AttachInternalIpRequest struct {
	InstanceId string `json:"instance_id"`

	Body *AttachInternalIpRequestBody `json:"body,omitempty"`
}

func (o AttachInternalIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachInternalIpRequest struct{}"
	}

	return strings.Join([]string{"AttachInternalIpRequest", string(data)}, " ")
}
