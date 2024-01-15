package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListLogStreamRequest struct {
	LogGroupId string `json:"log_group_id"`

	ContentType string `json:"Content-Type"`
}

func (o ListLogStreamRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogStreamRequest struct{}"
	}

	return strings.Join([]string{"ListLogStreamRequest", string(data)}, " ")
}
