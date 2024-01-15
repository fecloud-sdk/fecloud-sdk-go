package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type DeleteResourceTagRequest struct {
	Key string `json:"key"`

	ResourceId string `json:"resource_id"`

	ResourceType DeleteResourceTagRequestResourceType `json:"resource_type"`
}

func (o DeleteResourceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteResourceTagRequest", string(data)}, " ")
}

type DeleteResourceTagRequestResourceType struct {
	value string
}

type DeleteResourceTagRequestResourceTypeEnum struct {
	DC_DIRECTCONNECT DeleteResourceTagRequestResourceType
	DC_VGW           DeleteResourceTagRequestResourceType
	DC_VIF           DeleteResourceTagRequestResourceType
}

func GetDeleteResourceTagRequestResourceTypeEnum() DeleteResourceTagRequestResourceTypeEnum {
	return DeleteResourceTagRequestResourceTypeEnum{
		DC_DIRECTCONNECT: DeleteResourceTagRequestResourceType{
			value: "dc-directconnect",
		},
		DC_VGW: DeleteResourceTagRequestResourceType{
			value: "dc-vgw",
		},
		DC_VIF: DeleteResourceTagRequestResourceType{
			value: "dc-vif",
		},
	}
}

func (c DeleteResourceTagRequestResourceType) Value() string {
	return c.value
}

func (c DeleteResourceTagRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteResourceTagRequestResourceType) UnmarshalJSON(b []byte) error {
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
