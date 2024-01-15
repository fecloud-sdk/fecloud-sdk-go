package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListEngineProductsRequest struct {
	Engine ListEngineProductsRequestEngine `json:"engine"`

	ProductId *string `json:"product_id,omitempty"`
}

func (o ListEngineProductsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEngineProductsRequest struct{}"
	}

	return strings.Join([]string{"ListEngineProductsRequest", string(data)}, " ")
}

type ListEngineProductsRequestEngine struct {
	value string
}

type ListEngineProductsRequestEngineEnum struct {
	KAFKA ListEngineProductsRequestEngine
}

func GetListEngineProductsRequestEngineEnum() ListEngineProductsRequestEngineEnum {
	return ListEngineProductsRequestEngineEnum{
		KAFKA: ListEngineProductsRequestEngine{
			value: "kafka",
		},
	}
}

func (c ListEngineProductsRequestEngine) Value() string {
	return c.value
}

func (c ListEngineProductsRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEngineProductsRequestEngine) UnmarshalJSON(b []byte) error {
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
