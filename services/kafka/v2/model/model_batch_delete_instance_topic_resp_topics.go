package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteInstanceTopicRespTopics struct {
	Id *string `json:"id,omitempty"`

	Success *bool `json:"success,omitempty"`
}

func (o BatchDeleteInstanceTopicRespTopics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceTopicRespTopics struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceTopicRespTopics", string(data)}, " ")
}
