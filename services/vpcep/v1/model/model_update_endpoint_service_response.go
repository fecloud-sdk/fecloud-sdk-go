package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type UpdateEndpointServiceResponse struct {
	Id *string `json:"id,omitempty"`

	PortId *string `json:"port_id,omitempty"`

	ServiceName *string `json:"service_name,omitempty"`

	ServerType *UpdateEndpointServiceResponseServerType `json:"server_type,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	PoolId *string `json:"pool_id,omitempty"`

	ApprovalEnabled *bool `json:"approval_enabled,omitempty"`

	Status *UpdateEndpointServiceResponseStatus `json:"status,omitempty"`

	ServiceType *string `json:"service_type,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	Ports *[]PortList `json:"ports,omitempty"`

	TcpProxy *UpdateEndpointServiceResponseTcpProxy `json:"tcp_proxy,omitempty"`

	Tags *[]TagList `json:"tags,omitempty"`

	Description *string `json:"description,omitempty"`

	EnablePolicy   *bool `json:"enable_policy,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdateEndpointServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointServiceResponse struct{}"
	}

	return strings.Join([]string{"UpdateEndpointServiceResponse", string(data)}, " ")
}

type UpdateEndpointServiceResponseServerType struct {
	value string
}

type UpdateEndpointServiceResponseServerTypeEnum struct {
	VM  UpdateEndpointServiceResponseServerType
	VIP UpdateEndpointServiceResponseServerType
	LB  UpdateEndpointServiceResponseServerType
}

func GetUpdateEndpointServiceResponseServerTypeEnum() UpdateEndpointServiceResponseServerTypeEnum {
	return UpdateEndpointServiceResponseServerTypeEnum{
		VM: UpdateEndpointServiceResponseServerType{
			value: "VM",
		},
		VIP: UpdateEndpointServiceResponseServerType{
			value: "VIP",
		},
		LB: UpdateEndpointServiceResponseServerType{
			value: "LB",
		},
	}
}

func (c UpdateEndpointServiceResponseServerType) Value() string {
	return c.value
}

func (c UpdateEndpointServiceResponseServerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEndpointServiceResponseServerType) UnmarshalJSON(b []byte) error {
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

type UpdateEndpointServiceResponseStatus struct {
	value string
}

type UpdateEndpointServiceResponseStatusEnum struct {
	CREATING  UpdateEndpointServiceResponseStatus
	AVAILABLE UpdateEndpointServiceResponseStatus
	FAILED    UpdateEndpointServiceResponseStatus
}

func GetUpdateEndpointServiceResponseStatusEnum() UpdateEndpointServiceResponseStatusEnum {
	return UpdateEndpointServiceResponseStatusEnum{
		CREATING: UpdateEndpointServiceResponseStatus{
			value: "creating",
		},
		AVAILABLE: UpdateEndpointServiceResponseStatus{
			value: "available",
		},
		FAILED: UpdateEndpointServiceResponseStatus{
			value: "failed",
		},
	}
}

func (c UpdateEndpointServiceResponseStatus) Value() string {
	return c.value
}

func (c UpdateEndpointServiceResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEndpointServiceResponseStatus) UnmarshalJSON(b []byte) error {
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

type UpdateEndpointServiceResponseTcpProxy struct {
	value string
}

type UpdateEndpointServiceResponseTcpProxyEnum struct {
	CLOSE      UpdateEndpointServiceResponseTcpProxy
	TOA_OPEN   UpdateEndpointServiceResponseTcpProxy
	PROXY_OPEN UpdateEndpointServiceResponseTcpProxy
	OPEN       UpdateEndpointServiceResponseTcpProxy
	PROXY_VNI  UpdateEndpointServiceResponseTcpProxy
}

func GetUpdateEndpointServiceResponseTcpProxyEnum() UpdateEndpointServiceResponseTcpProxyEnum {
	return UpdateEndpointServiceResponseTcpProxyEnum{
		CLOSE: UpdateEndpointServiceResponseTcpProxy{
			value: "close",
		},
		TOA_OPEN: UpdateEndpointServiceResponseTcpProxy{
			value: "toa_open",
		},
		PROXY_OPEN: UpdateEndpointServiceResponseTcpProxy{
			value: "proxy_open",
		},
		OPEN: UpdateEndpointServiceResponseTcpProxy{
			value: "open",
		},
		PROXY_VNI: UpdateEndpointServiceResponseTcpProxy{
			value: "proxy_vni",
		},
	}
}

func (c UpdateEndpointServiceResponseTcpProxy) Value() string {
	return c.value
}

func (c UpdateEndpointServiceResponseTcpProxy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEndpointServiceResponseTcpProxy) UnmarshalJSON(b []byte) error {
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
