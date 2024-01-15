package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeShareNameRequest struct {
	ContentType string `json:"Content-Type"`

	ShareId string `json:"share_id"`

	Body *ChangeShareNameReq `json:"body,omitempty"`
}

func (o ChangeShareNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeShareNameRequest struct{}"
	}

	return strings.Join([]string{"ChangeShareNameRequest", string(data)}, " ")
}
