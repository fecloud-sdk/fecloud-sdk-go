package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateSecurityGroupRequestBody struct {
	SecurityGroupId string `json:"security_group_id"`
}

func (o UpdateSecurityGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSecurityGroupRequestBody", string(data)}, " ")
}
