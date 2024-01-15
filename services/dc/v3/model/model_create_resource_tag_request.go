package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateResourceTagRequest struct {
	ResourceId string `json:"resource_id"`

	ResourceType CreateResourceTagRequestResourceType `json:"resource_type"`

	Body *CreateResourceTagRequestBody `json:"body,omitempty"`
}

func (o CreateResourceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceTagRequest struct{}"
	}

	return strings.Join([]string{"CreateResourceTagRequest", string(data)}, " ")
}

type CreateResourceTagRequestResourceType struct {
	value string
}

type CreateResourceTagRequestResourceTypeEnum struct {
	DC_DIRECTCONNECT CreateResourceTagRequestResourceType
	DC_VGW           CreateResourceTagRequestResourceType
	DC_VIF           CreateResourceTagRequestResourceType
}

func GetCreateResourceTagRequestResourceTypeEnum() CreateResourceTagRequestResourceTypeEnum {
	return CreateResourceTagRequestResourceTypeEnum{
		DC_DIRECTCONNECT: CreateResourceTagRequestResourceType{
			value: "dc-directconnect",
		},
		DC_VGW: CreateResourceTagRequestResourceType{
			value: "dc-vgw",
		},
		DC_VIF: CreateResourceTagRequestResourceType{
			value: "dc-vif",
		},
	}
}

func (c CreateResourceTagRequestResourceType) Value() string {
	return c.value
}

func (c CreateResourceTagRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateResourceTagRequestResourceType) UnmarshalJSON(b []byte) error {
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
