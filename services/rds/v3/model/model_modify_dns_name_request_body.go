package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ModifyDnsNameRequestBody struct {
	DnsName string `json:"dns_name"`
}

func (o ModifyDnsNameRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyDnsNameRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyDnsNameRequestBody", string(data)}, " ")
}
