package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ContainerCidr struct {
	Cidr string `json:"cidr"`
}

func (o ContainerCidr) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerCidr struct{}"
	}

	return strings.Join([]string{"ContainerCidr", string(data)}, " ")
}
