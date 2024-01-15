package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListInstancesByTagsRequestBody struct {
	Offset *string `json:"offset,omitempty"`

	Limit *string `json:"limit,omitempty"`

	Action ListInstancesByTagsRequestBodyAction `json:"action"`

	Matches *[]MatchOption `json:"matches,omitempty"`

	Tags *[]TagOption `json:"tags,omitempty"`
}

func (o ListInstancesByTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesByTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ListInstancesByTagsRequestBody", string(data)}, " ")
}

type ListInstancesByTagsRequestBodyAction struct {
	value string
}

type ListInstancesByTagsRequestBodyActionEnum struct {
	FILTER ListInstancesByTagsRequestBodyAction
	COUNT  ListInstancesByTagsRequestBodyAction
}

func GetListInstancesByTagsRequestBodyActionEnum() ListInstancesByTagsRequestBodyActionEnum {
	return ListInstancesByTagsRequestBodyActionEnum{
		FILTER: ListInstancesByTagsRequestBodyAction{
			value: "filter",
		},
		COUNT: ListInstancesByTagsRequestBodyAction{
			value: "count",
		},
	}
}

func (c ListInstancesByTagsRequestBodyAction) Value() string {
	return c.value
}

func (c ListInstancesByTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesByTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
