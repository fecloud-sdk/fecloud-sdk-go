package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ExpandShareRequest struct {
	ContentType string `json:"Content-Type"`

	ShareId string `json:"share_id"`

	Body *ExpandShareRequestBody `json:"body,omitempty"`
}

func (o ExpandShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandShareRequest struct{}"
	}

	return strings.Join([]string{"ExpandShareRequest", string(data)}, " ")
}
