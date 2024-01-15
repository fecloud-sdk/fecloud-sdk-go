package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateIpRequest struct {
	InstanceId string `json:"instance_id"`

	Body *CreateIpRequestBody `json:"body,omitempty"`
}

func (o CreateIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpRequest struct{}"
	}

	return strings.Join([]string{"CreateIpRequest", string(data)}, " ")
}
