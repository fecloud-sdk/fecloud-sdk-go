package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeSecurityGroupRequestBody struct {
	ChangeSecurityGroup *ChangeSecurityGroup `json:"change_security_group"`
}

func (o ChangeSecurityGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSecurityGroupRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeSecurityGroupRequestBody", string(data)}, " ")
}
