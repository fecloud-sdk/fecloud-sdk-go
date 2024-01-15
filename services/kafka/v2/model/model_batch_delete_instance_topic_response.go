package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteInstanceTopicResponse struct {
	Topics         *[]BatchDeleteInstanceTopicRespTopics `json:"topics,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o BatchDeleteInstanceTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceTopicResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceTopicResponse", string(data)}, " ")
}
