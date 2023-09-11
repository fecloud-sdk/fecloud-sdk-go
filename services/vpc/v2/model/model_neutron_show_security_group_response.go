package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NeutronShowSecurityGroupResponse struct {
	SecurityGroup  *NeutronSecurityGroup `json:"security_group,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o NeutronShowSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"NeutronShowSecurityGroupResponse", string(data)}, " ")
}
