package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateSqlJobResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	JobId *string `json:"job_id,omitempty"`

	JobType *CreateSqlJobResponseJobType `json:"job_type,omitempty"`

	Schema *[]interface{} `json:"schema,omitempty"`

	Rows *[][]string `json:"rows,omitempty"`

	JobMode        *string `json:"job_mode,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSqlJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlJobResponse struct{}"
	}

	return strings.Join([]string{"CreateSqlJobResponse", string(data)}, " ")
}

type CreateSqlJobResponseJobType struct {
	value string
}

type CreateSqlJobResponseJobTypeEnum struct {
	DDL    CreateSqlJobResponseJobType
	DCL    CreateSqlJobResponseJobType
	IMPORT CreateSqlJobResponseJobType
	EXPORT CreateSqlJobResponseJobType
	QUERY  CreateSqlJobResponseJobType
	INSERT CreateSqlJobResponseJobType
}

func GetCreateSqlJobResponseJobTypeEnum() CreateSqlJobResponseJobTypeEnum {
	return CreateSqlJobResponseJobTypeEnum{
		DDL: CreateSqlJobResponseJobType{
			value: "DDL",
		},
		DCL: CreateSqlJobResponseJobType{
			value: "DCL",
		},
		IMPORT: CreateSqlJobResponseJobType{
			value: "IMPORT",
		},
		EXPORT: CreateSqlJobResponseJobType{
			value: "EXPORT",
		},
		QUERY: CreateSqlJobResponseJobType{
			value: "QUERY",
		},
		INSERT: CreateSqlJobResponseJobType{
			value: "INSERT",
		},
	}
}

func (c CreateSqlJobResponseJobType) Value() string {
	return c.value
}

func (c CreateSqlJobResponseJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSqlJobResponseJobType) UnmarshalJSON(b []byte) error {
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
