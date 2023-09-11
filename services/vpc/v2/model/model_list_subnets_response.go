package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSubnetsResponse struct {
	Subnets        *[]Subnet `json:"subnets,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSubnetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubnetsResponse struct{}"
	}

	return strings.Join([]string{"ListSubnetsResponse", string(data)}, " ")
}
