package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListResourceReq struct {
	Tags *[]TagWithMultiValue `json:"tags,omitempty"`

	TagsAny *[]TagWithMultiValue `json:"tags_any,omitempty"`

	NotTags *[]TagWithMultiValue `json:"not_tags,omitempty"`

	NotTagsAny *[]TagWithMultiValue `json:"not_tags_any,omitempty"`

	Action ListResourceReqAction `json:"action"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Matches *[]Match `json:"matches,omitempty"`
}

func (o ListResourceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceReq struct{}"
	}

	return strings.Join([]string{"ListResourceReq", string(data)}, " ")
}

type ListResourceReqAction struct {
	value string
}

type ListResourceReqActionEnum struct {
	FILTER ListResourceReqAction
	COUNT  ListResourceReqAction
}

func GetListResourceReqActionEnum() ListResourceReqActionEnum {
	return ListResourceReqActionEnum{
		FILTER: ListResourceReqAction{
			value: "filter",
		},
		COUNT: ListResourceReqAction{
			value: "count",
		},
	}
}

func (c ListResourceReqAction) Value() string {
	return c.value
}

func (c ListResourceReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourceReqAction) UnmarshalJSON(b []byte) error {
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
