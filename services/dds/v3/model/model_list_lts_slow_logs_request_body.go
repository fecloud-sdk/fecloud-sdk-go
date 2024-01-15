package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListLtsSlowLogsRequestBody struct {
	StartTime string `json:"start_time"`

	EndTime string `json:"end_time"`

	Limit int32 `json:"limit"`

	LineNum *string `json:"line_num,omitempty"`

	OperateType *ListLtsSlowLogsRequestBodyOperateType `json:"operate_type,omitempty"`

	NodeId *string `json:"node_id,omitempty"`

	Keywords *[]string `json:"keywords,omitempty"`

	DatabaseKeywords *[]string `json:"database_keywords,omitempty"`

	CollectionKeywords *[]string `json:"collection_keywords,omitempty"`

	MaxCostTime *int32 `json:"max_cost_time,omitempty"`

	MinCostTime *int32 `json:"min_cost_time,omitempty"`
}

func (o ListLtsSlowLogsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsSlowLogsRequestBody struct{}"
	}

	return strings.Join([]string{"ListLtsSlowLogsRequestBody", string(data)}, " ")
}

type ListLtsSlowLogsRequestBodyOperateType struct {
	value string
}

type ListLtsSlowLogsRequestBodyOperateTypeEnum struct {
	INSERT      ListLtsSlowLogsRequestBodyOperateType
	QUERY       ListLtsSlowLogsRequestBodyOperateType
	UPDATE      ListLtsSlowLogsRequestBodyOperateType
	REMOVE      ListLtsSlowLogsRequestBodyOperateType
	GETMORE     ListLtsSlowLogsRequestBodyOperateType
	COMMAND     ListLtsSlowLogsRequestBodyOperateType
	KILLCURSORS ListLtsSlowLogsRequestBodyOperateType
}

func GetListLtsSlowLogsRequestBodyOperateTypeEnum() ListLtsSlowLogsRequestBodyOperateTypeEnum {
	return ListLtsSlowLogsRequestBodyOperateTypeEnum{
		INSERT: ListLtsSlowLogsRequestBodyOperateType{
			value: "insert",
		},
		QUERY: ListLtsSlowLogsRequestBodyOperateType{
			value: "query",
		},
		UPDATE: ListLtsSlowLogsRequestBodyOperateType{
			value: "update",
		},
		REMOVE: ListLtsSlowLogsRequestBodyOperateType{
			value: "remove",
		},
		GETMORE: ListLtsSlowLogsRequestBodyOperateType{
			value: "getmore",
		},
		COMMAND: ListLtsSlowLogsRequestBodyOperateType{
			value: "command",
		},
		KILLCURSORS: ListLtsSlowLogsRequestBodyOperateType{
			value: "killcursors",
		},
	}
}

func (c ListLtsSlowLogsRequestBodyOperateType) Value() string {
	return c.value
}

func (c ListLtsSlowLogsRequestBodyOperateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLtsSlowLogsRequestBodyOperateType) UnmarshalJSON(b []byte) error {
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
