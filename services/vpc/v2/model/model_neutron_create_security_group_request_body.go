package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NeutronCreateSecurityGroupRequestBody struct {
	SecurityGroup *NeutronCreateSecurityGroupOption `json:"security_group"`
}

func (o NeutronCreateSecurityGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateSecurityGroupRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronCreateSecurityGroupRequestBody", string(data)}, " ")
}
