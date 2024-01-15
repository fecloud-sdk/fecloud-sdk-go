package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ServiceList struct {
	Id *string `json:"id,omitempty"`

	PortId *string `json:"port_id,omitempty"`

	ServiceName *string `json:"service_name,omitempty"`

	ServerType *string `json:"server_type,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	ApprovalEnabled *bool `json:"approval_enabled,omitempty"`

	Status *ServiceListStatus `json:"status,omitempty"`

	ServiceType *ServiceListServiceType `json:"service_type,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	DomainId *string `json:"domain_id,omitempty"`

	Ports *[]PortList `json:"ports,omitempty"`

	Tags *[]TagList `json:"tags,omitempty"`

	ConnectionCount *int32 `json:"connection_count,omitempty"`

	TcpProxy *ServiceListTcpProxy `json:"tcp_proxy,omitempty"`

	Error *[]Error `json:"error,omitempty"`

	Description *string `json:"description,omitempty"`

	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	EnablePolicy *bool `json:"enable_policy,omitempty"`
}

func (o ServiceList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceList struct{}"
	}

	return strings.Join([]string{"ServiceList", string(data)}, " ")
}

type ServiceListStatus struct {
	value string
}

type ServiceListStatusEnum struct {
	CREATING  ServiceListStatus
	AVAILABLE ServiceListStatus
	FAILED    ServiceListStatus
}

func GetServiceListStatusEnum() ServiceListStatusEnum {
	return ServiceListStatusEnum{
		CREATING: ServiceListStatus{
			value: "creating",
		},
		AVAILABLE: ServiceListStatus{
			value: "available",
		},
		FAILED: ServiceListStatus{
			value: "failed",
		},
	}
}

func (c ServiceListStatus) Value() string {
	return c.value
}

func (c ServiceListStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServiceListStatus) UnmarshalJSON(b []byte) error {
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

type ServiceListServiceType struct {
	value string
}

type ServiceListServiceTypeEnum struct {
	VM  ServiceListServiceType
	VIP ServiceListServiceType
	LB  ServiceListServiceType
}

func GetServiceListServiceTypeEnum() ServiceListServiceTypeEnum {
	return ServiceListServiceTypeEnum{
		VM: ServiceListServiceType{
			value: "VM",
		},
		VIP: ServiceListServiceType{
			value: "VIP",
		},
		LB: ServiceListServiceType{
			value: "LB",
		},
	}
}

func (c ServiceListServiceType) Value() string {
	return c.value
}

func (c ServiceListServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServiceListServiceType) UnmarshalJSON(b []byte) error {
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

type ServiceListTcpProxy struct {
	value string
}

type ServiceListTcpProxyEnum struct {
	CLOSE      ServiceListTcpProxy
	TOA_OPEN   ServiceListTcpProxy
	PROXY_OPEN ServiceListTcpProxy
	OPEN       ServiceListTcpProxy
	PROXY_VNI  ServiceListTcpProxy
}

func GetServiceListTcpProxyEnum() ServiceListTcpProxyEnum {
	return ServiceListTcpProxyEnum{
		CLOSE: ServiceListTcpProxy{
			value: "close",
		},
		TOA_OPEN: ServiceListTcpProxy{
			value: "toa_open",
		},
		PROXY_OPEN: ServiceListTcpProxy{
			value: "proxy_open",
		},
		OPEN: ServiceListTcpProxy{
			value: "open",
		},
		PROXY_VNI: ServiceListTcpProxy{
			value: "proxy_vni",
		},
	}
}

func (c ServiceListTcpProxy) Value() string {
	return c.value
}

func (c ServiceListTcpProxy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServiceListTcpProxy) UnmarshalJSON(b []byte) error {
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
