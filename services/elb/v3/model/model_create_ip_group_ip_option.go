package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateIpGroupIpOption struct {
	Ip string `json:"ip"`

	Description *string `json:"description,omitempty"`
}

func (o CreateIpGroupIpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpGroupIpOption struct{}"
	}

	return strings.Join([]string{"CreateIpGroupIpOption", string(data)}, " ")
}
