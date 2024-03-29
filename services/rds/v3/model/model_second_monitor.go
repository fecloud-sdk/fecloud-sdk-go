package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type SecondMonitor struct {
	SwitchOption bool `json:"switch_option"`

	Interval *SecondMonitorInterval `json:"interval,omitempty"`
}

func (o SecondMonitor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecondMonitor struct{}"
	}

	return strings.Join([]string{"SecondMonitor", string(data)}, " ")
}

type SecondMonitorInterval struct {
	value int32
}

type SecondMonitorIntervalEnum struct {
	E_1 SecondMonitorInterval
	E_5 SecondMonitorInterval
}

func GetSecondMonitorIntervalEnum() SecondMonitorIntervalEnum {
	return SecondMonitorIntervalEnum{
		E_1: SecondMonitorInterval{
			value: 1,
		}, E_5: SecondMonitorInterval{
			value: 5,
		},
	}
}

func (c SecondMonitorInterval) Value() int32 {
	return c.value
}

func (c SecondMonitorInterval) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SecondMonitorInterval) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
