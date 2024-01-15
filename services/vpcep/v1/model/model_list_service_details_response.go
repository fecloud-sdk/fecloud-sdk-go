package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListServiceDetailsResponse struct {
	Id *string `json:"id,omitempty"`

	PortId *string `json:"port_id,omitempty"`

	ServiceName *string `json:"service_name,omitempty"`

	ServerType *ListServiceDetailsResponseServerType `json:"server_type,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	ApprovalEnabled *bool `json:"approval_enabled,omitempty"`

	Status *ListServiceDetailsResponseStatus `json:"status,omitempty"`

	ServiceType *ListServiceDetailsResponseServiceType `json:"service_type,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	CidrType *string `json:"cidr_type,omitempty"`

	Ports *[]PortList `json:"ports,omitempty"`

	TcpProxy *ListServiceDetailsResponseTcpProxy `json:"tcp_proxy,omitempty"`

	Tags *[]TagList `json:"tags,omitempty"`

	Error *[]Error `json:"error,omitempty"`

	EnablePolicy *bool `json:"enable_policy,omitempty"`

	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListServiceDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListServiceDetailsResponse", string(data)}, " ")
}

type ListServiceDetailsResponseServerType struct {
	value string
}

type ListServiceDetailsResponseServerTypeEnum struct {
	VM  ListServiceDetailsResponseServerType
	VIP ListServiceDetailsResponseServerType
	LB  ListServiceDetailsResponseServerType
}

func GetListServiceDetailsResponseServerTypeEnum() ListServiceDetailsResponseServerTypeEnum {
	return ListServiceDetailsResponseServerTypeEnum{
		VM: ListServiceDetailsResponseServerType{
			value: "VM",
		},
		VIP: ListServiceDetailsResponseServerType{
			value: "VIP",
		},
		LB: ListServiceDetailsResponseServerType{
			value: "LB",
		},
	}
}

func (c ListServiceDetailsResponseServerType) Value() string {
	return c.value
}

func (c ListServiceDetailsResponseServerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServiceDetailsResponseServerType) UnmarshalJSON(b []byte) error {
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

type ListServiceDetailsResponseStatus struct {
	value string
}

type ListServiceDetailsResponseStatusEnum struct {
	CREATING  ListServiceDetailsResponseStatus
	AVAILABLE ListServiceDetailsResponseStatus
	FAILED    ListServiceDetailsResponseStatus
}

func GetListServiceDetailsResponseStatusEnum() ListServiceDetailsResponseStatusEnum {
	return ListServiceDetailsResponseStatusEnum{
		CREATING: ListServiceDetailsResponseStatus{
			value: "creating",
		},
		AVAILABLE: ListServiceDetailsResponseStatus{
			value: "available",
		},
		FAILED: ListServiceDetailsResponseStatus{
			value: "failed",
		},
	}
}

func (c ListServiceDetailsResponseStatus) Value() string {
	return c.value
}

func (c ListServiceDetailsResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServiceDetailsResponseStatus) UnmarshalJSON(b []byte) error {
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

type ListServiceDetailsResponseServiceType struct {
	value string
}

type ListServiceDetailsResponseServiceTypeEnum struct {
	GATEWAY   ListServiceDetailsResponseServiceType
	INTERFACE ListServiceDetailsResponseServiceType
}

func GetListServiceDetailsResponseServiceTypeEnum() ListServiceDetailsResponseServiceTypeEnum {
	return ListServiceDetailsResponseServiceTypeEnum{
		GATEWAY: ListServiceDetailsResponseServiceType{
			value: "gateway",
		},
		INTERFACE: ListServiceDetailsResponseServiceType{
			value: "interface",
		},
	}
}

func (c ListServiceDetailsResponseServiceType) Value() string {
	return c.value
}

func (c ListServiceDetailsResponseServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServiceDetailsResponseServiceType) UnmarshalJSON(b []byte) error {
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

type ListServiceDetailsResponseTcpProxy struct {
	value string
}

type ListServiceDetailsResponseTcpProxyEnum struct {
	CLOSE      ListServiceDetailsResponseTcpProxy
	TOA_OPEN   ListServiceDetailsResponseTcpProxy
	PROXY_OPEN ListServiceDetailsResponseTcpProxy
	OPEN       ListServiceDetailsResponseTcpProxy
	PROXY_VNI  ListServiceDetailsResponseTcpProxy
}

func GetListServiceDetailsResponseTcpProxyEnum() ListServiceDetailsResponseTcpProxyEnum {
	return ListServiceDetailsResponseTcpProxyEnum{
		CLOSE: ListServiceDetailsResponseTcpProxy{
			value: "close",
		},
		TOA_OPEN: ListServiceDetailsResponseTcpProxy{
			value: "toa_open",
		},
		PROXY_OPEN: ListServiceDetailsResponseTcpProxy{
			value: "proxy_open",
		},
		OPEN: ListServiceDetailsResponseTcpProxy{
			value: "open",
		},
		PROXY_VNI: ListServiceDetailsResponseTcpProxy{
			value: "proxy_vni",
		},
	}
}

func (c ListServiceDetailsResponseTcpProxy) Value() string {
	return c.value
}

func (c ListServiceDetailsResponseTcpProxy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServiceDetailsResponseTcpProxy) UnmarshalJSON(b []byte) error {
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
