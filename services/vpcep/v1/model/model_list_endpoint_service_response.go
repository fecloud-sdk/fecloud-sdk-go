package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListEndpointServiceResponse struct {
	EndpointServices *[]ServiceList `json:"endpoint_services,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEndpointServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointServiceResponse struct{}"
	}

	return strings.Join([]string{"ListEndpointServiceResponse", string(data)}, " ")
}
