package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AttachEipRequestBody struct {
	PublicIpId string `json:"public_ip_id"`

	PublicIp string `json:"public_ip"`
}

func (o AttachEipRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachEipRequestBody struct{}"
	}

	return strings.Join([]string{"AttachEipRequestBody", string(data)}, " ")
}
