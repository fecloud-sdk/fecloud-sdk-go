package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteConsumerGroupReq struct {
	Groups *[]string `json:"groups,omitempty"`
}

func (o BatchDeleteConsumerGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteConsumerGroupReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteConsumerGroupReq", string(data)}, " ")
}
