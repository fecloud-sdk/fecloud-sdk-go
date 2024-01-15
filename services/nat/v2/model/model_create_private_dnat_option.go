package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreatePrivateDnatOption struct {
	Description *string `json:"description,omitempty"`

	TransitIpId string `json:"transit_ip_id"`

	NetworkInterfaceId *string `json:"network_interface_id,omitempty"`

	GatewayId string `json:"gateway_id"`

	Protocol *CreatePrivateDnatOptionProtocol `json:"protocol,omitempty"`

	PrivateIpAddress *string `json:"private_ip_address,omitempty"`

	InternalServicePort *string `json:"internal_service_port,omitempty"`

	TransitServicePort *string `json:"transit_service_port,omitempty"`
}

func (o CreatePrivateDnatOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateDnatOption struct{}"
	}

	return strings.Join([]string{"CreatePrivateDnatOption", string(data)}, " ")
}

type CreatePrivateDnatOptionProtocol struct {
	value string
}

type CreatePrivateDnatOptionProtocolEnum struct {
	TCP CreatePrivateDnatOptionProtocol
	UDP CreatePrivateDnatOptionProtocol
	ANY CreatePrivateDnatOptionProtocol
}

func GetCreatePrivateDnatOptionProtocolEnum() CreatePrivateDnatOptionProtocolEnum {
	return CreatePrivateDnatOptionProtocolEnum{
		TCP: CreatePrivateDnatOptionProtocol{
			value: "tcp",
		},
		UDP: CreatePrivateDnatOptionProtocol{
			value: "udp",
		},
		ANY: CreatePrivateDnatOptionProtocol{
			value: "any",
		},
	}
}

func (c CreatePrivateDnatOptionProtocol) Value() string {
	return c.value
}

func (c CreatePrivateDnatOptionProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePrivateDnatOptionProtocol) UnmarshalJSON(b []byte) error {
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
