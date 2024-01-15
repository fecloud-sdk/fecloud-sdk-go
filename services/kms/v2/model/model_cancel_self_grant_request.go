package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CancelSelfGrantRequest struct {
	Body *RevokeGrantRequestBody `json:"body,omitempty"`
}

func (o CancelSelfGrantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSelfGrantRequest struct{}"
	}

	return strings.Join([]string{"CancelSelfGrantRequest", string(data)}, " ")
}
