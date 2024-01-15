package model

import (
	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/sdktime"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"
	"strings"
)

type VirtualInterface struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Bandwidth *int32 `json:"bandwidth,omitempty"`

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	Description *string `json:"description,omitempty"`

	DirectConnectId *string `json:"direct_connect_id,omitempty"`

	ServiceType *VirtualInterfaceServiceType `json:"service_type,omitempty"`

	Status *string `json:"status,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`

	Type *VirtualInterfaceType `json:"type,omitempty"`

	VgwId *string `json:"vgw_id,omitempty"`

	Vlan *int32 `json:"vlan,omitempty"`

	RouteLimit *int32 `json:"route_limit,omitempty"`

	EnableNqa *bool `json:"enable_nqa,omitempty"`

	EnableBfd *bool `json:"enable_bfd,omitempty"`

	LagId *string `json:"lag_id,omitempty"`

	DeviceId *string `json:"device_id,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Tags *[]Tag `json:"tags,omitempty"`

	VifPeers *[]VifPeer `json:"vif_peers,omitempty"`

	ExtendAttribute *VifExtendAttribute `json:"extend_attribute,omitempty"`
}

func (o VirtualInterface) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VirtualInterface struct{}"
	}

	return strings.Join([]string{"VirtualInterface", string(data)}, " ")
}

type VirtualInterfaceServiceType struct {
	value string
}

type VirtualInterfaceServiceTypeEnum struct {
	VGW  VirtualInterfaceServiceType
	GDGW VirtualInterfaceServiceType
	LGW  VirtualInterfaceServiceType
}

func GetVirtualInterfaceServiceTypeEnum() VirtualInterfaceServiceTypeEnum {
	return VirtualInterfaceServiceTypeEnum{
		VGW: VirtualInterfaceServiceType{
			value: "VGW",
		},
		GDGW: VirtualInterfaceServiceType{
			value: "GDGW",
		},
		LGW: VirtualInterfaceServiceType{
			value: "LGW",
		},
	}
}

func (c VirtualInterfaceServiceType) Value() string {
	return c.value
}

func (c VirtualInterfaceServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VirtualInterfaceServiceType) UnmarshalJSON(b []byte) error {
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

type VirtualInterfaceType struct {
	value string
}

type VirtualInterfaceTypeEnum struct {
	PRIVATE VirtualInterfaceType
	PUBLIC  VirtualInterfaceType
}

func GetVirtualInterfaceTypeEnum() VirtualInterfaceTypeEnum {
	return VirtualInterfaceTypeEnum{
		PRIVATE: VirtualInterfaceType{
			value: "private",
		},
		PUBLIC: VirtualInterfaceType{
			value: "public",
		},
	}
}

func (c VirtualInterfaceType) Value() string {
	return c.value
}

func (c VirtualInterfaceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VirtualInterfaceType) UnmarshalJSON(b []byte) error {
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
