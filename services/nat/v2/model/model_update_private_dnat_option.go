package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type UpdatePrivateDnatOption struct {
	Description *string `json:"description,omitempty"`

	TransitIpId *string `json:"transit_ip_id,omitempty"`

	NetworkInterfaceId *string `json:"network_interface_id,omitempty"`

	PrivateIpAddress *string `json:"private_ip_address,omitempty"`

	Protocol *UpdatePrivateDnatOptionProtocol `json:"protocol,omitempty"`

	InternalServicePort *string `json:"internal_service_port,omitempty"`

	TransitServicePort *string `json:"transit_service_port,omitempty"`
}

func (o UpdatePrivateDnatOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateDnatOption struct{}"
	}

	return strings.Join([]string{"UpdatePrivateDnatOption", string(data)}, " ")
}

type UpdatePrivateDnatOptionProtocol struct {
	value string
}

type UpdatePrivateDnatOptionProtocolEnum struct {
	TCP UpdatePrivateDnatOptionProtocol
	UDP UpdatePrivateDnatOptionProtocol
	ANY UpdatePrivateDnatOptionProtocol
}

func GetUpdatePrivateDnatOptionProtocolEnum() UpdatePrivateDnatOptionProtocolEnum {
	return UpdatePrivateDnatOptionProtocolEnum{
		TCP: UpdatePrivateDnatOptionProtocol{
			value: "tcp",
		},
		UDP: UpdatePrivateDnatOptionProtocol{
			value: "udp",
		},
		ANY: UpdatePrivateDnatOptionProtocol{
			value: "any",
		},
	}
}

func (c UpdatePrivateDnatOptionProtocol) Value() string {
	return c.value
}

func (c UpdatePrivateDnatOptionProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePrivateDnatOptionProtocol) UnmarshalJSON(b []byte) error {
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
