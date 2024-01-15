package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateVirtualInterface struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	DirectConnectId *string `json:"direct_connect_id,omitempty"`

	Type CreateVirtualInterfaceType `json:"type"`

	ServiceType *CreateVirtualInterfaceServiceType `json:"service_type,omitempty"`

	Vlan int32 `json:"vlan"`

	Bandwidth int32 `json:"bandwidth"`

	LocalGatewayV4Ip *string `json:"local_gateway_v4_ip,omitempty"`

	RemoteGatewayV4Ip *string `json:"remote_gateway_v4_ip,omitempty"`

	AddressFamily *string `json:"address_family,omitempty"`

	LocalGatewayV6Ip *string `json:"local_gateway_v6_ip,omitempty"`

	RemoteGatewayV6Ip *string `json:"remote_gateway_v6_ip,omitempty"`

	VgwId string `json:"vgw_id"`

	RouteMode CreateVirtualInterfaceRouteMode `json:"route_mode"`

	BgpAsn *int32 `json:"bgp_asn,omitempty"`

	BgpMd5 *string `json:"bgp_md5,omitempty"`

	RemoteEpGroup []string `json:"remote_ep_group"`

	ServiceEpGroup *[]string `json:"service_ep_group,omitempty"`

	EnableBfd *bool `json:"enable_bfd,omitempty"`

	EnableNqa *bool `json:"enable_nqa,omitempty"`

	LagId *string `json:"lag_id,omitempty"`

	ResourceTenantId *string `json:"resource_tenant_id,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateVirtualInterface) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirtualInterface struct{}"
	}

	return strings.Join([]string{"CreateVirtualInterface", string(data)}, " ")
}

type CreateVirtualInterfaceType struct {
	value string
}

type CreateVirtualInterfaceTypeEnum struct {
	PRIVATE CreateVirtualInterfaceType
	PUBLIC  CreateVirtualInterfaceType
}

func GetCreateVirtualInterfaceTypeEnum() CreateVirtualInterfaceTypeEnum {
	return CreateVirtualInterfaceTypeEnum{
		PRIVATE: CreateVirtualInterfaceType{
			value: "private",
		},
		PUBLIC: CreateVirtualInterfaceType{
			value: "public",
		},
	}
}

func (c CreateVirtualInterfaceType) Value() string {
	return c.value
}

func (c CreateVirtualInterfaceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVirtualInterfaceType) UnmarshalJSON(b []byte) error {
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

type CreateVirtualInterfaceServiceType struct {
	value string
}

type CreateVirtualInterfaceServiceTypeEnum struct {
	VPC  CreateVirtualInterfaceServiceType
	VGW  CreateVirtualInterfaceServiceType
	GDGW CreateVirtualInterfaceServiceType
	LGW  CreateVirtualInterfaceServiceType
}

func GetCreateVirtualInterfaceServiceTypeEnum() CreateVirtualInterfaceServiceTypeEnum {
	return CreateVirtualInterfaceServiceTypeEnum{
		VPC: CreateVirtualInterfaceServiceType{
			value: "vpc",
		},
		VGW: CreateVirtualInterfaceServiceType{
			value: "VGW",
		},
		GDGW: CreateVirtualInterfaceServiceType{
			value: "GDGW",
		},
		LGW: CreateVirtualInterfaceServiceType{
			value: "LGW",
		},
	}
}

func (c CreateVirtualInterfaceServiceType) Value() string {
	return c.value
}

func (c CreateVirtualInterfaceServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVirtualInterfaceServiceType) UnmarshalJSON(b []byte) error {
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

type CreateVirtualInterfaceRouteMode struct {
	value string
}

type CreateVirtualInterfaceRouteModeEnum struct {
	STATIC CreateVirtualInterfaceRouteMode
	BGP    CreateVirtualInterfaceRouteMode
}

func GetCreateVirtualInterfaceRouteModeEnum() CreateVirtualInterfaceRouteModeEnum {
	return CreateVirtualInterfaceRouteModeEnum{
		STATIC: CreateVirtualInterfaceRouteMode{
			value: "static",
		},
		BGP: CreateVirtualInterfaceRouteMode{
			value: "bgp",
		},
	}
}

func (c CreateVirtualInterfaceRouteMode) Value() string {
	return c.value
}

func (c CreateVirtualInterfaceRouteMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVirtualInterfaceRouteMode) UnmarshalJSON(b []byte) error {
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
