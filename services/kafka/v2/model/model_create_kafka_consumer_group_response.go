package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateKafkaConsumerGroupResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateKafkaConsumerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKafkaConsumerGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateKafkaConsumerGroupResponse", string(data)}, " ")
}
