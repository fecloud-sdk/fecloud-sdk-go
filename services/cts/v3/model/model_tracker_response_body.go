package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type TrackerResponseBody struct {
	Id *string `json:"id,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	Lts *Lts `json:"lts,omitempty"`

	TrackerType *TrackerResponseBodyTrackerType `json:"tracker_type,omitempty"`

	DomainId *string `json:"domain_id,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	TrackerName *string `json:"tracker_name,omitempty"`

	Status *TrackerResponseBodyStatus `json:"status,omitempty"`

	Detail *string `json:"detail,omitempty"`

	GroupId *string `json:"group_id,omitempty"`

	StreamId *string `json:"stream_id,omitempty"`

	ObsInfo *ObsInfo `json:"obs_info,omitempty"`
}

func (o TrackerResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrackerResponseBody struct{}"
	}

	return strings.Join([]string{"TrackerResponseBody", string(data)}, " ")
}

type TrackerResponseBodyTrackerType struct {
	value string
}

type TrackerResponseBodyTrackerTypeEnum struct {
	SYSTEM TrackerResponseBodyTrackerType
}

func GetTrackerResponseBodyTrackerTypeEnum() TrackerResponseBodyTrackerTypeEnum {
	return TrackerResponseBodyTrackerTypeEnum{
		SYSTEM: TrackerResponseBodyTrackerType{
			value: "system",
		},
	}
}

func (c TrackerResponseBodyTrackerType) Value() string {
	return c.value
}

func (c TrackerResponseBodyTrackerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TrackerResponseBodyTrackerType) UnmarshalJSON(b []byte) error {
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

type TrackerResponseBodyStatus struct {
	value string
}

type TrackerResponseBodyStatusEnum struct {
	ENABLED  TrackerResponseBodyStatus
	DISABLED TrackerResponseBodyStatus
}

func GetTrackerResponseBodyStatusEnum() TrackerResponseBodyStatusEnum {
	return TrackerResponseBodyStatusEnum{
		ENABLED: TrackerResponseBodyStatus{
			value: "enabled",
		},
		DISABLED: TrackerResponseBodyStatus{
			value: "disabled",
		},
	}
}

func (c TrackerResponseBodyStatus) Value() string {
	return c.value
}

func (c TrackerResponseBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TrackerResponseBodyStatus) UnmarshalJSON(b []byte) error {
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
