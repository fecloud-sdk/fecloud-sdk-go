package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListLogsRequestBody struct {
	StartTime string `json:"start_time"`

	EndTime string `json:"end_time"`

	Labels map[string]string `json:"labels,omitempty"`

	Keywords *string `json:"keywords,omitempty"`

	LineNum *string `json:"line_num,omitempty"`

	IsDesc *bool `json:"is_desc,omitempty"`

	SearchType *ListLogsRequestBodySearchType `json:"search_type,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Highlight *bool `json:"highlight,omitempty"`

	IsCount *bool `json:"is_count,omitempty"`

	IsIterative *bool `json:"is_iterative,omitempty"`
}

func (o ListLogsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogsRequestBody struct{}"
	}

	return strings.Join([]string{"ListLogsRequestBody", string(data)}, " ")
}

type ListLogsRequestBodySearchType struct {
	value string
}

type ListLogsRequestBodySearchTypeEnum struct {
	FORWARDS ListLogsRequestBodySearchType
}

func GetListLogsRequestBodySearchTypeEnum() ListLogsRequestBodySearchTypeEnum {
	return ListLogsRequestBodySearchTypeEnum{
		FORWARDS: ListLogsRequestBodySearchType{
			value: "forwards",
		},
	}
}

func (c ListLogsRequestBodySearchType) Value() string {
	return c.value
}

func (c ListLogsRequestBodySearchType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLogsRequestBodySearchType) UnmarshalJSON(b []byte) error {
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
