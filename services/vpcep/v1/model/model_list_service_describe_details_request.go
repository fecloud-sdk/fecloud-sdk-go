package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListServiceDescribeDetailsRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	EndpointServiceName *string `json:"endpoint_service_name,omitempty"`

	Id *string `json:"id,omitempty"`
}

func (o ListServiceDescribeDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceDescribeDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListServiceDescribeDetailsRequest", string(data)}, " ")
}
