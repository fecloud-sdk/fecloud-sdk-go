package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpadateIpGroupIpOption struct {
	Ip string `json:"ip"`

	Description *string `json:"description,omitempty"`
}

func (o UpadateIpGroupIpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpadateIpGroupIpOption struct{}"
	}

	return strings.Join([]string{"UpadateIpGroupIpOption", string(data)}, " ")
}
