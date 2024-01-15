package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListQueuesResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	Queues         *[]QueueInfo `json:"queues,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListQueuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueuesResponse struct{}"
	}

	return strings.Join([]string{"ListQueuesResponse", string(data)}, " ")
}
