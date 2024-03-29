package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NeutronDeleteFirewallPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NeutronDeleteFirewallPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteFirewallPolicyResponse struct{}"
	}

	return strings.Join([]string{"NeutronDeleteFirewallPolicyResponse", string(data)}, " ")
}
