package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateLogStreamRequest struct {
	LogGroupId string `json:"log_group_id"`

	ContentType string `json:"Content-Type"`

	Body *CreateLogStreamParams `json:"body,omitempty"`
}

func (o CreateLogStreamRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogStreamRequest struct{}"
	}

	return strings.Join([]string{"CreateLogStreamRequest", string(data)}, " ")
}
