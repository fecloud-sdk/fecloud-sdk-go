package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListQueuesRequest struct {
	QueueType *string `json:"queue_type,omitempty"`

	Tags *string `json:"tags,omitempty"`

	WithChargeInfo *bool `json:"with-charge-info,omitempty"`

	WithPriv *bool `json:"with-priv,omitempty"`
}

func (o ListQueuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueuesRequest struct{}"
	}

	return strings.Join([]string{"ListQueuesRequest", string(data)}, " ")
}
