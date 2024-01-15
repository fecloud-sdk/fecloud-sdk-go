package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateEndpointServiceNameResponse struct {
	EndpointServiceName *string `json:"endpoint_service_name,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o UpdateEndpointServiceNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointServiceNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateEndpointServiceNameResponse", string(data)}, " ")
}
