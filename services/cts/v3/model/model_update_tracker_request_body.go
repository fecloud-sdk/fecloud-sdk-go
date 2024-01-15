package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type UpdateTrackerRequestBody struct {
	TrackerType UpdateTrackerRequestBodyTrackerType `json:"tracker_type"`

	TrackerName string `json:"tracker_name"`

	Status *UpdateTrackerRequestBodyStatus `json:"status,omitempty"`

	IsLtsEnabled *bool `json:"is_lts_enabled,omitempty"`

	ObsInfo *TrackerObsInfo `json:"obs_info,omitempty"`
}

func (o UpdateTrackerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTrackerRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTrackerRequestBody", string(data)}, " ")
}

type UpdateTrackerRequestBodyTrackerType struct {
	value string
}

type UpdateTrackerRequestBodyTrackerTypeEnum struct {
	SYSTEM UpdateTrackerRequestBodyTrackerType
}

func GetUpdateTrackerRequestBodyTrackerTypeEnum() UpdateTrackerRequestBodyTrackerTypeEnum {
	return UpdateTrackerRequestBodyTrackerTypeEnum{
		SYSTEM: UpdateTrackerRequestBodyTrackerType{
			value: "system",
		},
	}
}

func (c UpdateTrackerRequestBodyTrackerType) Value() string {
	return c.value
}

func (c UpdateTrackerRequestBodyTrackerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTrackerRequestBodyTrackerType) UnmarshalJSON(b []byte) error {
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

type UpdateTrackerRequestBodyStatus struct {
	value string
}

type UpdateTrackerRequestBodyStatusEnum struct {
	ENABLED  UpdateTrackerRequestBodyStatus
	DISABLED UpdateTrackerRequestBodyStatus
}

func GetUpdateTrackerRequestBodyStatusEnum() UpdateTrackerRequestBodyStatusEnum {
	return UpdateTrackerRequestBodyStatusEnum{
		ENABLED: UpdateTrackerRequestBodyStatus{
			value: "enabled",
		},
		DISABLED: UpdateTrackerRequestBodyStatus{
			value: "disabled",
		},
	}
}

func (c UpdateTrackerRequestBodyStatus) Value() string {
	return c.value
}

func (c UpdateTrackerRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTrackerRequestBodyStatus) UnmarshalJSON(b []byte) error {
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
