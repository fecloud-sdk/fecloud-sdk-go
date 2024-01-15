package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteEndpointResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEndpointResponse struct{}"
	}

	return strings.Join([]string{"DeleteEndpointResponse", string(data)}, " ")
}
