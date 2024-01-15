package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NodeItem struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Status string `json:"status"`

	Role string `json:"role"`

	PrivateIp string `json:"private_ip"`

	PublicIp string `json:"public_ip"`

	SpecCode string `json:"spec_code"`

	AvailabilityZone string `json:"availability_zone"`
}

func (o NodeItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeItem struct{}"
	}

	return strings.Join([]string{"NodeItem", string(data)}, " ")
}
