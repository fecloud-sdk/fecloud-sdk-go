package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateTrackerRequestBody struct {
	TrackerType CreateTrackerRequestBodyTrackerType `json:"tracker_type"`

	TrackerName string `json:"tracker_name"`

	IsLtsEnabled *bool `json:"is_lts_enabled,omitempty"`

	ObsInfo *TrackerObsInfo `json:"obs_info,omitempty"`
}

func (o CreateTrackerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrackerRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTrackerRequestBody", string(data)}, " ")
}

type CreateTrackerRequestBodyTrackerType struct {
	value string
}

type CreateTrackerRequestBodyTrackerTypeEnum struct {
	SYSTEM CreateTrackerRequestBodyTrackerType
}

func GetCreateTrackerRequestBodyTrackerTypeEnum() CreateTrackerRequestBodyTrackerTypeEnum {
	return CreateTrackerRequestBodyTrackerTypeEnum{
		SYSTEM: CreateTrackerRequestBodyTrackerType{
			value: "system",
		},
	}
}

func (c CreateTrackerRequestBodyTrackerType) Value() string {
	return c.value
}

func (c CreateTrackerRequestBodyTrackerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTrackerRequestBodyTrackerType) UnmarshalJSON(b []byte) error {
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
