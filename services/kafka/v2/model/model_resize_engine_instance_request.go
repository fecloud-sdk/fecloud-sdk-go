package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ResizeEngineInstanceRequest struct {
	Engine ResizeEngineInstanceRequestEngine `json:"engine"`

	InstanceId string `json:"instance_id"`

	Body *ResizeEngineInstanceReq `json:"body,omitempty"`
}

func (o ResizeEngineInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeEngineInstanceRequest struct{}"
	}

	return strings.Join([]string{"ResizeEngineInstanceRequest", string(data)}, " ")
}

type ResizeEngineInstanceRequestEngine struct {
	value string
}

type ResizeEngineInstanceRequestEngineEnum struct {
	KAFKA ResizeEngineInstanceRequestEngine
}

func GetResizeEngineInstanceRequestEngineEnum() ResizeEngineInstanceRequestEngineEnum {
	return ResizeEngineInstanceRequestEngineEnum{
		KAFKA: ResizeEngineInstanceRequestEngine{
			value: "kafka",
		},
	}
}

func (c ResizeEngineInstanceRequestEngine) Value() string {
	return c.value
}

func (c ResizeEngineInstanceRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResizeEngineInstanceRequestEngine) UnmarshalJSON(b []byte) error {
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
