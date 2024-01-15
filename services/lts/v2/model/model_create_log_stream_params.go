package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateLogStreamParams struct {
	LogStreamName string `json:"log_stream_name"`

	TtlInDays *int32 `json:"ttl_in_days,omitempty"`

	Tags *[]TagsBody `json:"tags,omitempty"`
}

func (o CreateLogStreamParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogStreamParams struct{}"
	}

	return strings.Join([]string{"CreateLogStreamParams", string(data)}, " ")
}
