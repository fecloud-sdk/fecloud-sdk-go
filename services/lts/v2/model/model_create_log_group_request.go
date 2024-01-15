package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateLogGroupRequest struct {
	ContentType string `json:"Content-Type"`

	Body *CreateLogGroupParams `json:"body,omitempty"`
}

func (o CreateLogGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateLogGroupRequest", string(data)}, " ")
}
