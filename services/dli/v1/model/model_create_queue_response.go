package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateQueueResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	QueueName      *string `json:"queue_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateQueueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQueueResponse struct{}"
	}

	return strings.Join([]string{"CreateQueueResponse", string(data)}, " ")
}
