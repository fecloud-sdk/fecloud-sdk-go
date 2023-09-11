package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// NovaDisassociateSecurityGroupResponse Response Object
type NovaDisassociateSecurityGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NovaDisassociateSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaDisassociateSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"NovaDisassociateSecurityGroupResponse", string(data)}, " ")
}
