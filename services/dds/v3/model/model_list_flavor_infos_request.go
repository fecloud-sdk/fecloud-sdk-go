package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListFlavorInfosRequest struct {
	EngineName *ListFlavorInfosRequestEngineName `json:"engine_name,omitempty"`

	EngineVersion *string `json:"engine_version,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListFlavorInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorInfosRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorInfosRequest", string(data)}, " ")
}

type ListFlavorInfosRequestEngineName struct {
	value string
}

type ListFlavorInfosRequestEngineNameEnum struct {
	DDS_COMMUNITY ListFlavorInfosRequestEngineName
	DDS_ENHANCED  ListFlavorInfosRequestEngineName
}

func GetListFlavorInfosRequestEngineNameEnum() ListFlavorInfosRequestEngineNameEnum {
	return ListFlavorInfosRequestEngineNameEnum{
		DDS_COMMUNITY: ListFlavorInfosRequestEngineName{
			value: "DDS-Community",
		},
		DDS_ENHANCED: ListFlavorInfosRequestEngineName{
			value: "DDS-Enhanced",
		},
	}
}

func (c ListFlavorInfosRequestEngineName) Value() string {
	return c.value
}

func (c ListFlavorInfosRequestEngineName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlavorInfosRequestEngineName) UnmarshalJSON(b []byte) error {
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
