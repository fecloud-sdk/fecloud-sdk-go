package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteInstanceTopicReq struct {
	Topics *[]string `json:"topics,omitempty"`
}

func (o BatchDeleteInstanceTopicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceTopicReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceTopicReq", string(data)}, " ")
}
