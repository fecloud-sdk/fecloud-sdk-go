package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateInstanceTopicRequest struct {
	InstanceId string `json:"instance_id"`

	Body *CreateInstanceTopicReq `json:"body,omitempty"`
}

func (o CreateInstanceTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceTopicRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceTopicRequest", string(data)}, " ")
}
