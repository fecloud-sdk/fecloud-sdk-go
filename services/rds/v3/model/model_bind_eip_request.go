package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BindEipRequest struct {
	PublicIp *string `json:"public_ip,omitempty"`

	PublicIpId *string `json:"public_ip_id,omitempty"`

	IsBind bool `json:"is_bind"`
}

func (o BindEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindEipRequest struct{}"
	}

	return strings.Join([]string{"BindEipRequest", string(data)}, " ")
}
