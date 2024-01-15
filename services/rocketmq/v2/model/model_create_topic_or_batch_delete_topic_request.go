package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateTopicOrBatchDeleteTopicRequest struct {
	InstanceId string `json:"instance_id"`

	Action *string `json:"action,omitempty"`

	Body *CreateTopicOrBatchDeleteTopicReq `json:"body,omitempty"`
}

func (o CreateTopicOrBatchDeleteTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTopicOrBatchDeleteTopicRequest struct{}"
	}

	return strings.Join([]string{"CreateTopicOrBatchDeleteTopicRequest", string(data)}, " ")
}
