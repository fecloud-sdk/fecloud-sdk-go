package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateTrackerResponse struct {
	Id *string `json:"id,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	TrackerType *CreateTrackerResponseTrackerType `json:"tracker_type,omitempty"`

	DomainId *string `json:"domain_id,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	TrackerName *string `json:"tracker_name,omitempty"`

	Status *CreateTrackerResponseStatus `json:"status,omitempty"`

	Detail *string `json:"detail,omitempty"`

	ObsInfo        *ObsInfo `json:"obs_info,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o CreateTrackerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrackerResponse struct{}"
	}

	return strings.Join([]string{"CreateTrackerResponse", string(data)}, " ")
}

type CreateTrackerResponseTrackerType struct {
	value string
}

type CreateTrackerResponseTrackerTypeEnum struct {
	SYSTEM CreateTrackerResponseTrackerType
}

func GetCreateTrackerResponseTrackerTypeEnum() CreateTrackerResponseTrackerTypeEnum {
	return CreateTrackerResponseTrackerTypeEnum{
		SYSTEM: CreateTrackerResponseTrackerType{
			value: "system",
		},
	}
}

func (c CreateTrackerResponseTrackerType) Value() string {
	return c.value
}

func (c CreateTrackerResponseTrackerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTrackerResponseTrackerType) UnmarshalJSON(b []byte) error {
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

type CreateTrackerResponseStatus struct {
	value string
}

type CreateTrackerResponseStatusEnum struct {
	ENABLED  CreateTrackerResponseStatus
	DISABLED CreateTrackerResponseStatus
}

func GetCreateTrackerResponseStatusEnum() CreateTrackerResponseStatusEnum {
	return CreateTrackerResponseStatusEnum{
		ENABLED: CreateTrackerResponseStatus{
			value: "enabled",
		},
		DISABLED: CreateTrackerResponseStatus{
			value: "disabled",
		},
	}
}

func (c CreateTrackerResponseStatus) Value() string {
	return c.value
}

func (c CreateTrackerResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTrackerResponseStatus) UnmarshalJSON(b []byte) error {
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
