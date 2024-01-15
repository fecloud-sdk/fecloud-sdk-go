package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteConsumerGroupResp struct {
	JobId *string `json:"job_id,omitempty"`
}

func (o BatchDeleteConsumerGroupResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteConsumerGroupResp struct{}"
	}

	return strings.Join([]string{"BatchDeleteConsumerGroupResp", string(data)}, " ")
}
