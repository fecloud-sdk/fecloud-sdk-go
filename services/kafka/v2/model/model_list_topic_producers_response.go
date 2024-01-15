package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListTopicProducersResponse struct {
	Total *int32 `json:"total,omitempty"`

	Producers      *[]KafkaTopicProducerResponseProducers `json:"producers,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o ListTopicProducersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicProducersResponse struct{}"
	}

	return strings.Join([]string{"ListTopicProducersResponse", string(data)}, " ")
}
