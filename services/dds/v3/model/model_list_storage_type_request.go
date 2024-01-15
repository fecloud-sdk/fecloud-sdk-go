package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListStorageTypeRequest struct {
	EngineName *ListStorageTypeRequestEngineName `json:"engine_name,omitempty"`
}

func (o ListStorageTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStorageTypeRequest struct{}"
	}

	return strings.Join([]string{"ListStorageTypeRequest", string(data)}, " ")
}

type ListStorageTypeRequestEngineName struct {
	value string
}

type ListStorageTypeRequestEngineNameEnum struct {
	DDS_COMMUNITY ListStorageTypeRequestEngineName
	DDS_ENHANCED  ListStorageTypeRequestEngineName
}

func GetListStorageTypeRequestEngineNameEnum() ListStorageTypeRequestEngineNameEnum {
	return ListStorageTypeRequestEngineNameEnum{
		DDS_COMMUNITY: ListStorageTypeRequestEngineName{
			value: "DDS-Community",
		},
		DDS_ENHANCED: ListStorageTypeRequestEngineName{
			value: "DDS-Enhanced",
		},
	}
}

func (c ListStorageTypeRequestEngineName) Value() string {
	return c.value
}

func (c ListStorageTypeRequestEngineName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListStorageTypeRequestEngineName) UnmarshalJSON(b []byte) error {
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
