package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NovaDisassociateSecurityGroupRequest struct {
	ServerId string `json:"server_id"`

	Body *NovaDisassociateSecurityGroupRequestBody `json:"body,omitempty"`
}

func (o NovaDisassociateSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaDisassociateSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"NovaDisassociateSecurityGroupRequest", string(data)}, " ")
}
