package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PostPaidServerIpv6Bandwidth struct {
	Id *string `json:"id,omitempty"`
}

func (o PostPaidServerIpv6Bandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServerIpv6Bandwidth struct{}"
	}

	return strings.Join([]string{"PostPaidServerIpv6Bandwidth", string(data)}, " ")
}
