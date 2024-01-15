package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type LogStreamResBody struct {
	CreationTime *int64 `json:"creation_time,omitempty"`

	LogStreamId *string `json:"log_stream_id,omitempty"`

	LogStreamName *string `json:"log_stream_name,omitempty"`

	Tag map[string]string `json:"tag,omitempty"`

	FilterCount *int32 `json:"filter_count,omitempty"`

	IsFavorite *bool `json:"is_favorite,omitempty"`
}

func (o LogStreamResBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogStreamResBody struct{}"
	}

	return strings.Join([]string{"LogStreamResBody", string(data)}, " ")
}
