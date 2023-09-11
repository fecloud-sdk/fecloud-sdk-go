package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateSubnetResponse struct {
	Subnet         *SubnetResult `json:"subnet,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateSubnetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubnetResponse struct{}"
	}

	return strings.Join([]string{"UpdateSubnetResponse", string(data)}, " ")
}
