package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteLogStreamRequest struct {
	LogGroupId string `json:"log_group_id"`

	LogStreamId string `json:"log_stream_id"`

	ContentType string `json:"Content-Type"`
}

func (o DeleteLogStreamRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogStreamRequest struct{}"
	}

	return strings.Join([]string{"DeleteLogStreamRequest", string(data)}, " ")
}
