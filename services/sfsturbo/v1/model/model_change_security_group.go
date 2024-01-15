package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeSecurityGroup struct {
	SecurityGroupId string `json:"security_group_id"`
}

func (o ChangeSecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSecurityGroup struct{}"
	}

	return strings.Join([]string{"ChangeSecurityGroup", string(data)}, " ")
}
