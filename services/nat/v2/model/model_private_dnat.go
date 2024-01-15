package model

import (
	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/sdktime"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"
	"strings"
)

type PrivateDnat struct {
	Id *string `json:"id,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	Description *string `json:"description,omitempty"`

	TransitIpId *string `json:"transit_ip_id,omitempty"`

	GatewayId *string `json:"gateway_id,omitempty"`

	NetworkInterfaceId *string `json:"network_interface_id,omitempty"`

	Type *string `json:"type,omitempty"`

	Protocol *PrivateDnatProtocol `json:"protocol,omitempty"`

	PrivateIpAddress *string `json:"private_ip_address,omitempty"`

	InternalServicePort *string `json:"internal_service_port,omitempty"`

	TransitServicePort *string `json:"transit_service_port,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o PrivateDnat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateDnat struct{}"
	}

	return strings.Join([]string{"PrivateDnat", string(data)}, " ")
}

type PrivateDnatProtocol struct {
	value string
}

type PrivateDnatProtocolEnum struct {
	TCP PrivateDnatProtocol
	UDP PrivateDnatProtocol
	ANY PrivateDnatProtocol
}

func GetPrivateDnatProtocolEnum() PrivateDnatProtocolEnum {
	return PrivateDnatProtocolEnum{
		TCP: PrivateDnatProtocol{
			value: "tcp",
		},
		UDP: PrivateDnatProtocol{
			value: "udp",
		},
		ANY: PrivateDnatProtocol{
			value: "any",
		},
	}
}

func (c PrivateDnatProtocol) Value() string {
	return c.value
}

func (c PrivateDnatProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrivateDnatProtocol) UnmarshalJSON(b []byte) error {
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
