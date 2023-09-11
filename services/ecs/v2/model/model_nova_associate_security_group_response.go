package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// NovaAssociateSecurityGroupResponse Response Object
type NovaAssociateSecurityGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NovaAssociateSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaAssociateSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"NovaAssociateSecurityGroupResponse", string(data)}, " ")
}
