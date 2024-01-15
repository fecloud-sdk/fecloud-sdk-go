package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListServiceDetailsRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`
}

func (o ListServiceDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListServiceDetailsRequest", string(data)}, " ")
}
