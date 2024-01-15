package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateLogDumpObsRequestBody struct {
	LogGroupId string `json:"log_group_id"`

	LogStreamIds []string `json:"log_stream_ids"`

	ObsBucketName string `json:"obs_bucket_name"`

	Type string `json:"type"`

	StorageFormat string `json:"storage_format"`

	SwitchOn *bool `json:"switch_on,omitempty"`

	PrefixName *string `json:"prefix_name,omitempty"`

	DirPrefixName *string `json:"dir_prefix_name,omitempty"`

	Period CreateLogDumpObsRequestBodyPeriod `json:"period"`

	PeriodUnit CreateLogDumpObsRequestBodyPeriodUnit `json:"period_unit"`
}

func (o CreateLogDumpObsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogDumpObsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateLogDumpObsRequestBody", string(data)}, " ")
}

type CreateLogDumpObsRequestBodyPeriod struct {
	value int32
}

type CreateLogDumpObsRequestBodyPeriodEnum struct {
	E_1  CreateLogDumpObsRequestBodyPeriod
	E_2  CreateLogDumpObsRequestBodyPeriod
	E_3  CreateLogDumpObsRequestBodyPeriod
	E_5  CreateLogDumpObsRequestBodyPeriod
	E_6  CreateLogDumpObsRequestBodyPeriod
	E_12 CreateLogDumpObsRequestBodyPeriod
	E_30 CreateLogDumpObsRequestBodyPeriod
}

func GetCreateLogDumpObsRequestBodyPeriodEnum() CreateLogDumpObsRequestBodyPeriodEnum {
	return CreateLogDumpObsRequestBodyPeriodEnum{
		E_1: CreateLogDumpObsRequestBodyPeriod{
			value: 1,
		}, E_2: CreateLogDumpObsRequestBodyPeriod{
			value: 2,
		}, E_3: CreateLogDumpObsRequestBodyPeriod{
			value: 3,
		}, E_5: CreateLogDumpObsRequestBodyPeriod{
			value: 5,
		}, E_6: CreateLogDumpObsRequestBodyPeriod{
			value: 6,
		}, E_12: CreateLogDumpObsRequestBodyPeriod{
			value: 12,
		}, E_30: CreateLogDumpObsRequestBodyPeriod{
			value: 30,
		},
	}
}

func (c CreateLogDumpObsRequestBodyPeriod) Value() int32 {
	return c.value
}

func (c CreateLogDumpObsRequestBodyPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLogDumpObsRequestBodyPeriod) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type CreateLogDumpObsRequestBodyPeriodUnit struct {
	value string
}

type CreateLogDumpObsRequestBodyPeriodUnitEnum struct {
	MIN  CreateLogDumpObsRequestBodyPeriodUnit
	HOUR CreateLogDumpObsRequestBodyPeriodUnit
}

func GetCreateLogDumpObsRequestBodyPeriodUnitEnum() CreateLogDumpObsRequestBodyPeriodUnitEnum {
	return CreateLogDumpObsRequestBodyPeriodUnitEnum{
		MIN: CreateLogDumpObsRequestBodyPeriodUnit{
			value: "\"min\"",
		},
		HOUR: CreateLogDumpObsRequestBodyPeriodUnit{
			value: "\"hour\"",
		},
	}
}

func (c CreateLogDumpObsRequestBodyPeriodUnit) Value() string {
	return c.value
}

func (c CreateLogDumpObsRequestBodyPeriodUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLogDumpObsRequestBodyPeriodUnit) UnmarshalJSON(b []byte) error {
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
