package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateEndpointServiceRequestBody struct {
	PortId string `json:"port_id"`

	ServiceName *string `json:"service_name,omitempty"`

	VpcId string `json:"vpc_id"`

	ApprovalEnabled *bool `json:"approval_enabled,omitempty"`

	ServiceType *string `json:"service_type,omitempty"`

	ServerType CreateEndpointServiceRequestBodyServerType `json:"server_type"`

	Ports []PortList `json:"ports"`

	TcpProxy *CreateEndpointServiceRequestBodyTcpProxy `json:"tcp_proxy,omitempty"`

	Tags *[]TagList `json:"tags,omitempty"`

	Description *string `json:"description,omitempty"`

	EnablePolicy *bool `json:"enable_policy,omitempty"`
}

func (o CreateEndpointServiceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointServiceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEndpointServiceRequestBody", string(data)}, " ")
}

type CreateEndpointServiceRequestBodyServerType struct {
	value string
}

type CreateEndpointServiceRequestBodyServerTypeEnum struct {
	VM  CreateEndpointServiceRequestBodyServerType
	VIP CreateEndpointServiceRequestBodyServerType
	LB  CreateEndpointServiceRequestBodyServerType
}

func GetCreateEndpointServiceRequestBodyServerTypeEnum() CreateEndpointServiceRequestBodyServerTypeEnum {
	return CreateEndpointServiceRequestBodyServerTypeEnum{
		VM: CreateEndpointServiceRequestBodyServerType{
			value: "VM",
		},
		VIP: CreateEndpointServiceRequestBodyServerType{
			value: "VIP",
		},
		LB: CreateEndpointServiceRequestBodyServerType{
			value: "LB",
		},
	}
}

func (c CreateEndpointServiceRequestBodyServerType) Value() string {
	return c.value
}

func (c CreateEndpointServiceRequestBodyServerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEndpointServiceRequestBodyServerType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateEndpointServiceRequestBodyTcpProxy struct {
	value string
}

type CreateEndpointServiceRequestBodyTcpProxyEnum struct {
	CLOSE      CreateEndpointServiceRequestBodyTcpProxy
	TOA_OPEN   CreateEndpointServiceRequestBodyTcpProxy
	PROXY_OPEN CreateEndpointServiceRequestBodyTcpProxy
	OPEN       CreateEndpointServiceRequestBodyTcpProxy
	PROXY_VNI  CreateEndpointServiceRequestBodyTcpProxy
}

func GetCreateEndpointServiceRequestBodyTcpProxyEnum() CreateEndpointServiceRequestBodyTcpProxyEnum {
	return CreateEndpointServiceRequestBodyTcpProxyEnum{
		CLOSE: CreateEndpointServiceRequestBodyTcpProxy{
			value: "close",
		},
		TOA_OPEN: CreateEndpointServiceRequestBodyTcpProxy{
			value: "toa_open",
		},
		PROXY_OPEN: CreateEndpointServiceRequestBodyTcpProxy{
			value: "proxy_open",
		},
		OPEN: CreateEndpointServiceRequestBodyTcpProxy{
			value: "open",
		},
		PROXY_VNI: CreateEndpointServiceRequestBodyTcpProxy{
			value: "proxy_vni",
		},
	}
}

func (c CreateEndpointServiceRequestBodyTcpProxy) Value() string {
	return c.value
}

func (c CreateEndpointServiceRequestBodyTcpProxy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEndpointServiceRequestBodyTcpProxy) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
