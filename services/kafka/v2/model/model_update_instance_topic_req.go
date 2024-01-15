package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceTopicReq struct {
	Topics *[]UpdateInstanceTopicReqTopics `json:"topics,omitempty"`
}

func (o UpdateInstanceTopicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceTopicReq struct{}"
	}

	return strings.Join([]string{"UpdateInstanceTopicReq", string(data)}, " ")
}
