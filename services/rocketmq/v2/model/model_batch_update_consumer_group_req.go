package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchUpdateConsumerGroupReq struct {
	Groups *[]ConsumerGroup `json:"groups,omitempty"`
}

func (o BatchUpdateConsumerGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateConsumerGroupReq struct{}"
	}

	return strings.Join([]string{"BatchUpdateConsumerGroupReq", string(data)}, " ")
}
