package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ShowResourceTagRequest struct {
	ResourceType ShowResourceTagRequestResourceType `json:"resource_type"`

	ResourceId string `json:"resource_id"`
}

func (o ShowResourceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceTagRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceTagRequest", string(data)}, " ")
}

type ShowResourceTagRequestResourceType struct {
	value string
}

type ShowResourceTagRequestResourceTypeEnum struct {
	DC_DIRECTCONNECT ShowResourceTagRequestResourceType
	DC_VGW           ShowResourceTagRequestResourceType
	DC_VIF           ShowResourceTagRequestResourceType
}

func GetShowResourceTagRequestResourceTypeEnum() ShowResourceTagRequestResourceTypeEnum {
	return ShowResourceTagRequestResourceTypeEnum{
		DC_DIRECTCONNECT: ShowResourceTagRequestResourceType{
			value: "dc-directconnect",
		},
		DC_VGW: ShowResourceTagRequestResourceType{
			value: "dc-vgw",
		},
		DC_VIF: ShowResourceTagRequestResourceType{
			value: "dc-vif",
		},
	}
}

func (c ShowResourceTagRequestResourceType) Value() string {
	return c.value
}

func (c ShowResourceTagRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowResourceTagRequestResourceType) UnmarshalJSON(b []byte) error {
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
