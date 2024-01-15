package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListFlavorsRequest struct {
	Region *string `json:"region,omitempty"`

	EngineName *ListFlavorsRequestEngineName `json:"engine_name,omitempty"`
}

func (o ListFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorsRequest", string(data)}, " ")
}

type ListFlavorsRequestEngineName struct {
	value string
}

type ListFlavorsRequestEngineNameEnum struct {
	DDS_COMMUNITY ListFlavorsRequestEngineName
	DDS_ENHANCED  ListFlavorsRequestEngineName
}

func GetListFlavorsRequestEngineNameEnum() ListFlavorsRequestEngineNameEnum {
	return ListFlavorsRequestEngineNameEnum{
		DDS_COMMUNITY: ListFlavorsRequestEngineName{
			value: "DDS-Community",
		},
		DDS_ENHANCED: ListFlavorsRequestEngineName{
			value: "DDS-Enhanced",
		},
	}
}

func (c ListFlavorsRequestEngineName) Value() string {
	return c.value
}

func (c ListFlavorsRequestEngineName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlavorsRequestEngineName) UnmarshalJSON(b []byte) error {
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
