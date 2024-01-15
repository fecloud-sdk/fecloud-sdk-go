package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListSqlJobsRequest struct {
	CurrentPage *int32 `json:"current-page,omitempty"`

	DbName *string `json:"db_name,omitempty"`

	End *int64 `json:"end,omitempty"`

	EngineType *string `json:"engine-type,omitempty"`

	JobStatus *string `json:"job-status,omitempty"`

	JobType *ListSqlJobsRequestJobType `json:"job-type,omitempty"`

	Order *ListSqlJobsRequestOrder `json:"order,omitempty"`

	Owner *string `json:"owner,omitempty"`

	PageSize *int32 `json:"page-size,omitempty"`

	QueueName *string `json:"queue_name,omitempty"`

	SqlPattern *string `json:"sql_pattern,omitempty"`

	Start *int64 `json:"start,omitempty"`

	TableName *string `json:"table_name,omitempty"`

	Tags *string `json:"tags,omitempty"`
}

func (o ListSqlJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlJobsRequest struct{}"
	}

	return strings.Join([]string{"ListSqlJobsRequest", string(data)}, " ")
}

type ListSqlJobsRequestJobType struct {
	value string
}

type ListSqlJobsRequestJobTypeEnum struct {
	DDL    ListSqlJobsRequestJobType
	DCL    ListSqlJobsRequestJobType
	IMPORT ListSqlJobsRequestJobType
	EXPORT ListSqlJobsRequestJobType
	QUERY  ListSqlJobsRequestJobType
	INSERT ListSqlJobsRequestJobType
}

func GetListSqlJobsRequestJobTypeEnum() ListSqlJobsRequestJobTypeEnum {
	return ListSqlJobsRequestJobTypeEnum{
		DDL: ListSqlJobsRequestJobType{
			value: "DDL",
		},
		DCL: ListSqlJobsRequestJobType{
			value: "DCL",
		},
		IMPORT: ListSqlJobsRequestJobType{
			value: "IMPORT",
		},
		EXPORT: ListSqlJobsRequestJobType{
			value: "EXPORT",
		},
		QUERY: ListSqlJobsRequestJobType{
			value: "QUERY",
		},
		INSERT: ListSqlJobsRequestJobType{
			value: "INSERT",
		},
	}
}

func (c ListSqlJobsRequestJobType) Value() string {
	return c.value
}

func (c ListSqlJobsRequestJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSqlJobsRequestJobType) UnmarshalJSON(b []byte) error {
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

type ListSqlJobsRequestOrder struct {
	value string
}

type ListSqlJobsRequestOrderEnum struct {
	DURATION_DESC   ListSqlJobsRequestOrder
	DURATION_ASC    ListSqlJobsRequestOrder
	START_TIME_DESC ListSqlJobsRequestOrder
	START_TIME_ASC  ListSqlJobsRequestOrder
}

func GetListSqlJobsRequestOrderEnum() ListSqlJobsRequestOrderEnum {
	return ListSqlJobsRequestOrderEnum{
		DURATION_DESC: ListSqlJobsRequestOrder{
			value: "duration_desc",
		},
		DURATION_ASC: ListSqlJobsRequestOrder{
			value: "duration_asc",
		},
		START_TIME_DESC: ListSqlJobsRequestOrder{
			value: "start_time_desc",
		},
		START_TIME_ASC: ListSqlJobsRequestOrder{
			value: "start_time_asc",
		},
	}
}

func (c ListSqlJobsRequestOrder) Value() string {
	return c.value
}

func (c ListSqlJobsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSqlJobsRequestOrder) UnmarshalJSON(b []byte) error {
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
