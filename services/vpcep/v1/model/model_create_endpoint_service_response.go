package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateEndpointServiceResponse struct {
	Id *string `json:"id,omitempty"`

	PortId *string `json:"port_id,omitempty"`

	ServiceName *string `json:"service_name,omitempty"`

	ServerType *CreateEndpointServiceResponseServerType `json:"server_type,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	PoolId *string `json:"pool_id,omitempty"`

	ApprovalEnabled *bool `json:"approval_enabled,omitempty"`

	Status *CreateEndpointServiceResponseStatus `json:"status,omitempty"`

	ServiceType *string `json:"service_type,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	Ports *[]PortList `json:"ports,omitempty"`

	TcpProxy *CreateEndpointServiceResponseTcpProxy `json:"tcp_proxy,omitempty"`

	Tags *[]TagList `json:"tags,omitempty"`

	Description *string `json:"description,omitempty"`

	EnablePolicy   *bool `json:"enable_policy,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateEndpointServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointServiceResponse struct{}"
	}

	return strings.Join([]string{"CreateEndpointServiceResponse", string(data)}, " ")
}

type CreateEndpointServiceResponseServerType struct {
	value string
}

type CreateEndpointServiceResponseServerTypeEnum struct {
	VM  CreateEndpointServiceResponseServerType
	VIP CreateEndpointServiceResponseServerType
	LB  CreateEndpointServiceResponseServerType
}

func GetCreateEndpointServiceResponseServerTypeEnum() CreateEndpointServiceResponseServerTypeEnum {
	return CreateEndpointServiceResponseServerTypeEnum{
		VM: CreateEndpointServiceResponseServerType{
			value: "VM",
		},
		VIP: CreateEndpointServiceResponseServerType{
			value: "VIP",
		},
		LB: CreateEndpointServiceResponseServerType{
			value: "LB",
		},
	}
}

func (c CreateEndpointServiceResponseServerType) Value() string {
	return c.value
}

func (c CreateEndpointServiceResponseServerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEndpointServiceResponseServerType) UnmarshalJSON(b []byte) error {
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

type CreateEndpointServiceResponseStatus struct {
	value string
}

type CreateEndpointServiceResponseStatusEnum struct {
	CREATING  CreateEndpointServiceResponseStatus
	AVAILABLE CreateEndpointServiceResponseStatus
	FAILED    CreateEndpointServiceResponseStatus
}

func GetCreateEndpointServiceResponseStatusEnum() CreateEndpointServiceResponseStatusEnum {
	return CreateEndpointServiceResponseStatusEnum{
		CREATING: CreateEndpointServiceResponseStatus{
			value: "creating",
		},
		AVAILABLE: CreateEndpointServiceResponseStatus{
			value: "available",
		},
		FAILED: CreateEndpointServiceResponseStatus{
			value: "failed",
		},
	}
}

func (c CreateEndpointServiceResponseStatus) Value() string {
	return c.value
}

func (c CreateEndpointServiceResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEndpointServiceResponseStatus) UnmarshalJSON(b []byte) error {
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

type CreateEndpointServiceResponseTcpProxy struct {
	value string
}

type CreateEndpointServiceResponseTcpProxyEnum struct {
	CLOSE      CreateEndpointServiceResponseTcpProxy
	TOA_OPEN   CreateEndpointServiceResponseTcpProxy
	PROXY_OPEN CreateEndpointServiceResponseTcpProxy
	OPEN       CreateEndpointServiceResponseTcpProxy
	PROXY_VNI  CreateEndpointServiceResponseTcpProxy
}

func GetCreateEndpointServiceResponseTcpProxyEnum() CreateEndpointServiceResponseTcpProxyEnum {
	return CreateEndpointServiceResponseTcpProxyEnum{
		CLOSE: CreateEndpointServiceResponseTcpProxy{
			value: "close",
		},
		TOA_OPEN: CreateEndpointServiceResponseTcpProxy{
			value: "toa_open",
		},
		PROXY_OPEN: CreateEndpointServiceResponseTcpProxy{
			value: "proxy_open",
		},
		OPEN: CreateEndpointServiceResponseTcpProxy{
			value: "open",
		},
		PROXY_VNI: CreateEndpointServiceResponseTcpProxy{
			value: "proxy_vni",
		},
	}
}

func (c CreateEndpointServiceResponseTcpProxy) Value() string {
	return c.value
}

func (c CreateEndpointServiceResponseTcpProxy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEndpointServiceResponseTcpProxy) UnmarshalJSON(b []byte) error {
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
