package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NeutronCreateSecurityGroupResponse struct {
	SecurityGroup  *NeutronSecurityGroup `json:"security_group,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o NeutronCreateSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"NeutronCreateSecurityGroupResponse", string(data)}, " ")
}
