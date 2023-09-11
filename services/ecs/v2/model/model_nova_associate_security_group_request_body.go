package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// NovaAssociateSecurityGroupRequestBody This is a auto create Body Object
type NovaAssociateSecurityGroupRequestBody struct {
	AddSecurityGroup *NovaAddSecurityGroupOption `json:"addSecurityGroup"`
}

func (o NovaAssociateSecurityGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaAssociateSecurityGroupRequestBody struct{}"
	}

	return strings.Join([]string{"NovaAssociateSecurityGroupRequestBody", string(data)}, " ")
}
