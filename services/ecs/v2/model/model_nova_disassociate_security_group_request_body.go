package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// NovaDisassociateSecurityGroupRequestBody This is a auto create Body Object
type NovaDisassociateSecurityGroupRequestBody struct {
	RemoveSecurityGroup *NovaRemoveSecurityGroupOption `json:"removeSecurityGroup"`
}

func (o NovaDisassociateSecurityGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaDisassociateSecurityGroupRequestBody struct{}"
	}

	return strings.Join([]string{"NovaDisassociateSecurityGroupRequestBody", string(data)}, " ")
}
