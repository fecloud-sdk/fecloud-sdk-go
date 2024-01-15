package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResetUserPasswrodRequest struct {
	InstanceId string `json:"instance_id"`

	UserName string `json:"user_name"`

	Body *ResetUserPasswrodReq `json:"body,omitempty"`
}

func (o ResetUserPasswrodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetUserPasswrodRequest struct{}"
	}

	return strings.Join([]string{"ResetUserPasswrodRequest", string(data)}, " ")
}
