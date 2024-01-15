package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ShowInstanceExtendProductInfoRequest struct {
	InstanceId string `json:"instance_id"`

	Type ShowInstanceExtendProductInfoRequestType `json:"type"`

	Engine ShowInstanceExtendProductInfoRequestEngine `json:"engine"`
}

func (o ShowInstanceExtendProductInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceExtendProductInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceExtendProductInfoRequest", string(data)}, " ")
}

type ShowInstanceExtendProductInfoRequestType struct {
	value string
}

type ShowInstanceExtendProductInfoRequestTypeEnum struct {
	ADVANCED ShowInstanceExtendProductInfoRequestType
	PLATINUM ShowInstanceExtendProductInfoRequestType
	DEC      ShowInstanceExtendProductInfoRequestType
	EXP      ShowInstanceExtendProductInfoRequestType
}

func GetShowInstanceExtendProductInfoRequestTypeEnum() ShowInstanceExtendProductInfoRequestTypeEnum {
	return ShowInstanceExtendProductInfoRequestTypeEnum{
		ADVANCED: ShowInstanceExtendProductInfoRequestType{
			value: "advanced",
		},
		PLATINUM: ShowInstanceExtendProductInfoRequestType{
			value: "platinum",
		},
		DEC: ShowInstanceExtendProductInfoRequestType{
			value: "dec",
		},
		EXP: ShowInstanceExtendProductInfoRequestType{
			value: "exp",
		},
	}
}

func (c ShowInstanceExtendProductInfoRequestType) Value() string {
	return c.value
}

func (c ShowInstanceExtendProductInfoRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceExtendProductInfoRequestType) UnmarshalJSON(b []byte) error {
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

type ShowInstanceExtendProductInfoRequestEngine struct {
	value string
}

type ShowInstanceExtendProductInfoRequestEngineEnum struct {
	KAFKA ShowInstanceExtendProductInfoRequestEngine
}

func GetShowInstanceExtendProductInfoRequestEngineEnum() ShowInstanceExtendProductInfoRequestEngineEnum {
	return ShowInstanceExtendProductInfoRequestEngineEnum{
		KAFKA: ShowInstanceExtendProductInfoRequestEngine{
			value: "kafka",
		},
	}
}

func (c ShowInstanceExtendProductInfoRequestEngine) Value() string {
	return c.value
}

func (c ShowInstanceExtendProductInfoRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceExtendProductInfoRequestEngine) UnmarshalJSON(b []byte) error {
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
