package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListTagResourceInstancesRequestBody struct {
	Offset *string `json:"offset,omitempty"`

	Limit *string `json:"limit,omitempty"`

	Action ListTagResourceInstancesRequestBodyAction `json:"action"`

	Matches *[]Match `json:"matches,omitempty"`

	NotTags *[]Tags `json:"not_tags,omitempty"`

	Tags *[]Tags `json:"tags,omitempty"`

	TagsAny *[]Tags `json:"tags_any,omitempty"`

	NotTagsAny *[]Tags `json:"not_tags_any,omitempty"`

	SysTags *[]Tags `json:"sys_tags,omitempty"`
}

func (o ListTagResourceInstancesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagResourceInstancesRequestBody struct{}"
	}

	return strings.Join([]string{"ListTagResourceInstancesRequestBody", string(data)}, " ")
}

type ListTagResourceInstancesRequestBodyAction struct {
	value string
}

type ListTagResourceInstancesRequestBodyActionEnum struct {
	FILTER ListTagResourceInstancesRequestBodyAction
	COUNT  ListTagResourceInstancesRequestBodyAction
}

func GetListTagResourceInstancesRequestBodyActionEnum() ListTagResourceInstancesRequestBodyActionEnum {
	return ListTagResourceInstancesRequestBodyActionEnum{
		FILTER: ListTagResourceInstancesRequestBodyAction{
			value: "filter",
		},
		COUNT: ListTagResourceInstancesRequestBodyAction{
			value: "count",
		},
	}
}

func (c ListTagResourceInstancesRequestBodyAction) Value() string {
	return c.value
}

func (c ListTagResourceInstancesRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagResourceInstancesRequestBodyAction) UnmarshalJSON(b []byte) error {
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
