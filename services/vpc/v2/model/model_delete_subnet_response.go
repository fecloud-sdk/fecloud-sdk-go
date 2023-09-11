package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteSubnetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSubnetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubnetResponse struct{}"
	}

	return strings.Join([]string{"DeleteSubnetResponse", string(data)}, " ")
}
