package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteTopicResp struct {
	JobId *string `json:"job_id,omitempty"`
}

func (o BatchDeleteTopicResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTopicResp struct{}"
	}

	return strings.Join([]string{"BatchDeleteTopicResp", string(data)}, " ")
}
