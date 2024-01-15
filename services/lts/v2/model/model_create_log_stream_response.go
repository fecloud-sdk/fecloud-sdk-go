package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateLogStreamResponse struct {
	LogStreamId    *string `json:"log_stream_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLogStreamResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogStreamResponse struct{}"
	}

	return strings.Join([]string{"CreateLogStreamResponse", string(data)}, " ")
}
