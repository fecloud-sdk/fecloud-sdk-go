package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NeutronDeleteFirewallGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NeutronDeleteFirewallGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteFirewallGroupResponse struct{}"
	}

	return strings.Join([]string{"NeutronDeleteFirewallGroupResponse", string(data)}, " ")
}
