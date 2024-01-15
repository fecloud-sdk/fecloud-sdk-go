package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateFlinkSqlJobTemplateRequestBody struct {
	Name string `json:"name"`

	Desc *string `json:"desc,omitempty"`

	SqlBody *string `json:"sql_body,omitempty"`

	Tags *[]TmsTag `json:"tags,omitempty"`

	JobType *CreateFlinkSqlJobTemplateRequestBodyJobType `json:"job_type,omitempty"`
}

func (o CreateFlinkSqlJobTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkSqlJobTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFlinkSqlJobTemplateRequestBody", string(data)}, " ")
}

type CreateFlinkSqlJobTemplateRequestBodyJobType struct {
	value string
}

type CreateFlinkSqlJobTemplateRequestBodyJobTypeEnum struct {
	FLINK_SQL_JOB            CreateFlinkSqlJobTemplateRequestBodyJobType
	FLINK_OPENSOURCE_SQL_JOB CreateFlinkSqlJobTemplateRequestBodyJobType
}

func GetCreateFlinkSqlJobTemplateRequestBodyJobTypeEnum() CreateFlinkSqlJobTemplateRequestBodyJobTypeEnum {
	return CreateFlinkSqlJobTemplateRequestBodyJobTypeEnum{
		FLINK_SQL_JOB: CreateFlinkSqlJobTemplateRequestBodyJobType{
			value: "flink_sql_job",
		},
		FLINK_OPENSOURCE_SQL_JOB: CreateFlinkSqlJobTemplateRequestBodyJobType{
			value: "flink_opensource_sql_job",
		},
	}
}

func (c CreateFlinkSqlJobTemplateRequestBodyJobType) Value() string {
	return c.value
}

func (c CreateFlinkSqlJobTemplateRequestBodyJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFlinkSqlJobTemplateRequestBodyJobType) UnmarshalJSON(b []byte) error {
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
