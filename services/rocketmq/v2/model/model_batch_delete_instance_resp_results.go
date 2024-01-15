package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type BatchDeleteInstanceRespResults struct {
	Result *BatchDeleteInstanceRespResultsResult `json:"result,omitempty"`

	Instance *string `json:"instance,omitempty"`
}

func (o BatchDeleteInstanceRespResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceRespResults struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceRespResults", string(data)}, " ")
}

type BatchDeleteInstanceRespResultsResult struct {
	value string
}

type BatchDeleteInstanceRespResultsResultEnum struct {
	SUCCESS BatchDeleteInstanceRespResultsResult
	FAILED  BatchDeleteInstanceRespResultsResult
}

func GetBatchDeleteInstanceRespResultsResultEnum() BatchDeleteInstanceRespResultsResultEnum {
	return BatchDeleteInstanceRespResultsResultEnum{
		SUCCESS: BatchDeleteInstanceRespResultsResult{
			value: "success",
		},
		FAILED: BatchDeleteInstanceRespResultsResult{
			value: "failed",
		},
	}
}

func (c BatchDeleteInstanceRespResultsResult) Value() string {
	return c.value
}

func (c BatchDeleteInstanceRespResultsResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteInstanceRespResultsResult) UnmarshalJSON(b []byte) error {
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
