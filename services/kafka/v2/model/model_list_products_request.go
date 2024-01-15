package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListProductsRequest struct {
	Engine ListProductsRequestEngine `json:"engine"`
}

func (o ListProductsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsRequest struct{}"
	}

	return strings.Join([]string{"ListProductsRequest", string(data)}, " ")
}

type ListProductsRequestEngine struct {
	value string
}

type ListProductsRequestEngineEnum struct {
	KAFKA ListProductsRequestEngine
}

func GetListProductsRequestEngineEnum() ListProductsRequestEngineEnum {
	return ListProductsRequestEngineEnum{
		KAFKA: ListProductsRequestEngine{
			value: "kafka",
		},
	}
}

func (c ListProductsRequestEngine) Value() string {
	return c.value
}

func (c ListProductsRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProductsRequestEngine) UnmarshalJSON(b []byte) error {
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
