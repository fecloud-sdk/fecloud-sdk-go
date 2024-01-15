package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeSecurityGroupRequest struct {
	ContentType string `json:"Content-Type"`

	ShareId string `json:"share_id"`

	Body *ChangeSecurityGroupRequestBody `json:"body,omitempty"`
}

func (o ChangeSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"ChangeSecurityGroupRequest", string(data)}, " ")
}
