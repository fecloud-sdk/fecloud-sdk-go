package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ShowEngineInstanceExtendProductInfoRequest struct {
	Engine ShowEngineInstanceExtendProductInfoRequestEngine `json:"engine"`

	InstanceId string `json:"instance_id"`

	Type ShowEngineInstanceExtendProductInfoRequestType `json:"type"`
}

func (o ShowEngineInstanceExtendProductInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEngineInstanceExtendProductInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowEngineInstanceExtendProductInfoRequest", string(data)}, " ")
}

type ShowEngineInstanceExtendProductInfoRequestEngine struct {
	value string
}

type ShowEngineInstanceExtendProductInfoRequestEngineEnum struct {
	KAFKA ShowEngineInstanceExtendProductInfoRequestEngine
}

func GetShowEngineInstanceExtendProductInfoRequestEngineEnum() ShowEngineInstanceExtendProductInfoRequestEngineEnum {
	return ShowEngineInstanceExtendProductInfoRequestEngineEnum{
		KAFKA: ShowEngineInstanceExtendProductInfoRequestEngine{
			value: "kafka",
		},
	}
}

func (c ShowEngineInstanceExtendProductInfoRequestEngine) Value() string {
	return c.value
}

func (c ShowEngineInstanceExtendProductInfoRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEngineInstanceExtendProductInfoRequestEngine) UnmarshalJSON(b []byte) error {
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

type ShowEngineInstanceExtendProductInfoRequestType struct {
	value string
}

type ShowEngineInstanceExtendProductInfoRequestTypeEnum struct {
	ADVANCED ShowEngineInstanceExtendProductInfoRequestType
}

func GetShowEngineInstanceExtendProductInfoRequestTypeEnum() ShowEngineInstanceExtendProductInfoRequestTypeEnum {
	return ShowEngineInstanceExtendProductInfoRequestTypeEnum{
		ADVANCED: ShowEngineInstanceExtendProductInfoRequestType{
			value: "advanced",
		},
	}
}

func (c ShowEngineInstanceExtendProductInfoRequestType) Value() string {
	return c.value
}

func (c ShowEngineInstanceExtendProductInfoRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEngineInstanceExtendProductInfoRequestType) UnmarshalJSON(b []byte) error {
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
