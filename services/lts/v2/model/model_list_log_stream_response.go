package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListLogStreamResponse struct {
	LogStreams     *[]LogStreamResBody `json:"log_streams,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListLogStreamResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogStreamResponse struct{}"
	}

	return strings.Join([]string{"ListLogStreamResponse", string(data)}, " ")
}
