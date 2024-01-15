package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateShareRequest struct {
	ContentType string `json:"Content-Type"`

	Body *CreateShareRequestBody `json:"body,omitempty"`
}

func (o CreateShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShareRequest struct{}"
	}

	return strings.Join([]string{"CreateShareRequest", string(data)}, " ")
}
