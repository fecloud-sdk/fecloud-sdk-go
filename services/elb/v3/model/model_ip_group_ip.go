package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type IpGroupIp struct {
	Ip string `json:"ip"`
}

func (o IpGroupIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpGroupIp struct{}"
	}

	return strings.Join([]string{"IpGroupIp", string(data)}, " ")
}
