package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchUpdateConsumerGroupResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchUpdateConsumerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateConsumerGroupResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateConsumerGroupResponse", string(data)}, " ")
}
