package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ShowSqlJobStatusResponse struct {
	JobId *string `json:"job_id,omitempty"`

	JobType *string `json:"job_type,omitempty"`

	QueueName *string `json:"queue_name,omitempty"`

	Owner *string `json:"owner,omitempty"`

	StartTime *int64 `json:"start_time,omitempty"`

	Duration *int64 `json:"duration,omitempty"`

	Status *ShowSqlJobStatusResponseStatus `json:"status,omitempty"`

	InputRowCount *int64 `json:"input_row_count,omitempty"`

	BadRowCount *int64 `json:"bad_row_count,omitempty"`

	InputSize *int64 `json:"input_size,omitempty"`

	ResultCount *int32 `json:"result_count,omitempty"`

	DatabaseName *string `json:"database_name,omitempty"`

	TableName *string `json:"table_name,omitempty"`

	Detail *string `json:"detail,omitempty"`

	UserConf *string `json:"user_conf,omitempty"`

	ResultPath *string `json:"result_path,omitempty"`

	ResultFormat *string `json:"result_format,omitempty"`

	Statement *string `json:"statement,omitempty"`

	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	JobMode *string `json:"job_mode,omitempty"`

	Tags           *[]TmsTag `json:"tags,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowSqlJobStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlJobStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlJobStatusResponse", string(data)}, " ")
}

type ShowSqlJobStatusResponseStatus struct {
	value string
}

type ShowSqlJobStatusResponseStatusEnum struct {
	LAUNCHING ShowSqlJobStatusResponseStatus
	RUNNING   ShowSqlJobStatusResponseStatus
	FINISHED  ShowSqlJobStatusResponseStatus
	FAILED    ShowSqlJobStatusResponseStatus
	CANCELLED ShowSqlJobStatusResponseStatus
}

func GetShowSqlJobStatusResponseStatusEnum() ShowSqlJobStatusResponseStatusEnum {
	return ShowSqlJobStatusResponseStatusEnum{
		LAUNCHING: ShowSqlJobStatusResponseStatus{
			value: "LAUNCHING",
		},
		RUNNING: ShowSqlJobStatusResponseStatus{
			value: "RUNNING",
		},
		FINISHED: ShowSqlJobStatusResponseStatus{
			value: "FINISHED",
		},
		FAILED: ShowSqlJobStatusResponseStatus{
			value: "FAILED",
		},
		CANCELLED: ShowSqlJobStatusResponseStatus{
			value: "CANCELLED",
		},
	}
}

func (c ShowSqlJobStatusResponseStatus) Value() string {
	return c.value
}

func (c ShowSqlJobStatusResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSqlJobStatusResponseStatus) UnmarshalJSON(b []byte) error {
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
