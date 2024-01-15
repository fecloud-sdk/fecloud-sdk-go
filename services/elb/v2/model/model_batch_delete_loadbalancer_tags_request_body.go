package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type BatchDeleteLoadbalancerTagsRequestBody struct {
	Action BatchDeleteLoadbalancerTagsRequestBodyAction `json:"action"`

	Tags []ResourceTag `json:"tags"`
}

func (o BatchDeleteLoadbalancerTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteLoadbalancerTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteLoadbalancerTagsRequestBody", string(data)}, " ")
}

type BatchDeleteLoadbalancerTagsRequestBodyAction struct {
	value string
}

type BatchDeleteLoadbalancerTagsRequestBodyActionEnum struct {
	DELETE BatchDeleteLoadbalancerTagsRequestBodyAction
}

func GetBatchDeleteLoadbalancerTagsRequestBodyActionEnum() BatchDeleteLoadbalancerTagsRequestBodyActionEnum {
	return BatchDeleteLoadbalancerTagsRequestBodyActionEnum{
		DELETE: BatchDeleteLoadbalancerTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchDeleteLoadbalancerTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchDeleteLoadbalancerTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteLoadbalancerTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
