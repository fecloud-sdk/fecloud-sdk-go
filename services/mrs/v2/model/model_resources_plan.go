package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ResourcesPlan struct {
	PeriodType string `json:"period_type"`

	StartTime string `json:"start_time"`

	EndTime string `json:"end_time"`

	MinCapacity int32 `json:"min_capacity"`

	MaxCapacity int32 `json:"max_capacity"`

	EffectiveDays *[]ResourcesPlanEffectiveDays `json:"effective_days,omitempty"`
}

func (o ResourcesPlan) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesPlan struct{}"
	}

	return strings.Join([]string{"ResourcesPlan", string(data)}, " ")
}

type ResourcesPlanEffectiveDays struct {
	value string
}

type ResourcesPlanEffectiveDaysEnum struct {
	MONDAY    ResourcesPlanEffectiveDays
	TUESDAY   ResourcesPlanEffectiveDays
	WEDNESDAY ResourcesPlanEffectiveDays
	THURSDAY  ResourcesPlanEffectiveDays
	FRIDAY    ResourcesPlanEffectiveDays
	SATURDAY  ResourcesPlanEffectiveDays
	SUNDAY    ResourcesPlanEffectiveDays
}

func GetResourcesPlanEffectiveDaysEnum() ResourcesPlanEffectiveDaysEnum {
	return ResourcesPlanEffectiveDaysEnum{
		MONDAY: ResourcesPlanEffectiveDays{
			value: "MONDAY",
		},
		TUESDAY: ResourcesPlanEffectiveDays{
			value: "TUESDAY",
		},
		WEDNESDAY: ResourcesPlanEffectiveDays{
			value: "WEDNESDAY",
		},
		THURSDAY: ResourcesPlanEffectiveDays{
			value: "THURSDAY",
		},
		FRIDAY: ResourcesPlanEffectiveDays{
			value: "FRIDAY",
		},
		SATURDAY: ResourcesPlanEffectiveDays{
			value: "SATURDAY",
		},
		SUNDAY: ResourcesPlanEffectiveDays{
			value: "SUNDAY",
		},
	}
}

func (c ResourcesPlanEffectiveDays) Value() string {
	return c.value
}

func (c ResourcesPlanEffectiveDays) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourcesPlanEffectiveDays) UnmarshalJSON(b []byte) error {
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
