package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListLogsRequest struct {
	LogGroupId string `json:"log_group_id"`

	LogStreamId string `json:"log_stream_id"`

	ContentType string `json:"Content-Type"`

	Body *ListLogsRequestBody `json:"body,omitempty"`
}

func (o ListLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogsRequest struct{}"
	}

	return strings.Join([]string{"ListLogsRequest", string(data)}, " ")
}
