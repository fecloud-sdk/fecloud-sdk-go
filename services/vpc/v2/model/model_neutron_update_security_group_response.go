package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NeutronUpdateSecurityGroupResponse struct {
	SecurityGroup  *NeutronSecurityGroup `json:"security_group,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o NeutronUpdateSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"NeutronUpdateSecurityGroupResponse", string(data)}, " ")
}
