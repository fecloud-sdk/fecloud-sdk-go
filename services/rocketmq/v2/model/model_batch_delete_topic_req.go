package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteTopicReq struct {
	Topics *[]string `json:"topics,omitempty"`
}

func (o BatchDeleteTopicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTopicReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteTopicReq", string(data)}, " ")
}
