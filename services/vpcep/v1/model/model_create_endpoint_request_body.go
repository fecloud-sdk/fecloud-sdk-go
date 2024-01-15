package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateEndpointRequestBody struct {
	SubnetId *string `json:"subnet_id,omitempty"`

	EndpointServiceId string `json:"endpoint_service_id"`

	VpcId string `json:"vpc_id"`

	EnableDns *bool `json:"enable_dns,omitempty"`

	Tags *[]TagList `json:"tags,omitempty"`

	Routetables *[]string `json:"routetables,omitempty"`

	PortIp *string `json:"port_ip,omitempty"`

	Whitelist *[]string `json:"whitelist,omitempty"`

	EnableWhitelist *bool `json:"enable_whitelist,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o CreateEndpointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEndpointRequestBody", string(data)}, " ")
}
