package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ShowJobResponse struct {
	Status *ShowJobResponseStatus `json:"status,omitempty"`

	Entities *JobEntities `json:"entities,omitempty"`

	JobId *string `json:"job_id,omitempty"`

	JobType *string `json:"job_type,omitempty"`

	BeginTime *string `json:"begin_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	FailReason     *string `json:"fail_reason,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobResponse struct{}"
	}

	return strings.Join([]string{"ShowJobResponse", string(data)}, " ")
}

type ShowJobResponseStatus struct {
	value string
}

type ShowJobResponseStatusEnum struct {
	SUCCESS ShowJobResponseStatus
	RUNNING ShowJobResponseStatus
	FAIL    ShowJobResponseStatus
	INIT    ShowJobResponseStatus
}

func GetShowJobResponseStatusEnum() ShowJobResponseStatusEnum {
	return ShowJobResponseStatusEnum{
		SUCCESS: ShowJobResponseStatus{
			value: "SUCCESS",
		},
		RUNNING: ShowJobResponseStatus{
			value: "RUNNING",
		},
		FAIL: ShowJobResponseStatus{
			value: "FAIL",
		},
		INIT: ShowJobResponseStatus{
			value: "INIT",
		},
	}
}

func (c ShowJobResponseStatus) Value() string {
	return c.value
}

func (c ShowJobResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobResponseStatus) UnmarshalJSON(b []byte) error {
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
