package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type PortList struct {
	ClientPort *int32 `json:"client_port,omitempty"`

	ServerPort *int32 `json:"server_port,omitempty"`

	Protocol *PortListProtocol `json:"protocol,omitempty"`
}

func (o PortList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortList struct{}"
	}

	return strings.Join([]string{"PortList", string(data)}, " ")
}

type PortListProtocol struct {
	value string
}

type PortListProtocolEnum struct {
	TCP PortListProtocol
}

func GetPortListProtocolEnum() PortListProtocolEnum {
	return PortListProtocolEnum{
		TCP: PortListProtocol{
			value: "TCP",
		},
	}
}

func (c PortListProtocol) Value() string {
	return c.value
}

func (c PortListProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PortListProtocol) UnmarshalJSON(b []byte) error {
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
