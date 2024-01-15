package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CancelGrantRequest struct {
	Body *RevokeGrantRequestBody `json:"body,omitempty"`
}

func (o CancelGrantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelGrantRequest struct{}"
	}

	return strings.Join([]string{"CancelGrantRequest", string(data)}, " ")
}
