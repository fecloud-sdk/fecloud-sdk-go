package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type SqlJobInfo struct {
	JobId string `json:"job_id"`

	JobType string `json:"job_type"`

	QueueName string `json:"queue_name"`

	Owner string `json:"owner"`

	StartTime int64 `json:"start_time"`

	Duration *int32 `json:"duration,omitempty"`

	Status SqlJobInfoStatus `json:"status"`

	InputRowCount *int64 `json:"input_row_count,omitempty"`

	BadRowCount *int64 `json:"bad_row_count,omitempty"`

	InputSize int64 `json:"input_size"`

	ResultCount int32 `json:"result_count"`

	DatabaseName *string `json:"database_name,omitempty"`

	TableName *string `json:"table_name,omitempty"`

	WithColumnHeader *bool `json:"with_column_header,omitempty"`

	Detail string `json:"detail"`

	Statement string `json:"statement"`

	Tags *[]TmsTag `json:"tags,omitempty"`

	Message *string `json:"message,omitempty"`

	EndTime *int64 `json:"end_time,omitempty"`

	CpuCost *string `json:"cpu_cost,omitempty"`

	OutputByte *string `json:"output_byte,omitempty"`
}

func (o SqlJobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlJobInfo struct{}"
	}

	return strings.Join([]string{"SqlJobInfo", string(data)}, " ")
}

type SqlJobInfoStatus struct {
	value string
}

type SqlJobInfoStatusEnum struct {
	LAUNCHING SqlJobInfoStatus
	RUNNING   SqlJobInfoStatus
	FINISHED  SqlJobInfoStatus
	FAILED    SqlJobInfoStatus
	CANCELLED SqlJobInfoStatus
}

func GetSqlJobInfoStatusEnum() SqlJobInfoStatusEnum {
	return SqlJobInfoStatusEnum{
		LAUNCHING: SqlJobInfoStatus{
			value: "LAUNCHING",
		},
		RUNNING: SqlJobInfoStatus{
			value: "RUNNING",
		},
		FINISHED: SqlJobInfoStatus{
			value: "FINISHED",
		},
		FAILED: SqlJobInfoStatus{
			value: "FAILED",
		},
		CANCELLED: SqlJobInfoStatus{
			value: "CANCELLED",
		},
	}
}

func (c SqlJobInfoStatus) Value() string {
	return c.value
}

func (c SqlJobInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlJobInfoStatus) UnmarshalJSON(b []byte) error {
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
