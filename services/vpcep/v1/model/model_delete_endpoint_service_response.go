package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteEndpointServiceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEndpointServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEndpointServiceResponse struct{}"
	}

	return strings.Join([]string{"DeleteEndpointServiceResponse", string(data)}, " ")
}
