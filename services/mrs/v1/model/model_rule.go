package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type Rule struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	AdjustmentType RuleAdjustmentType `json:"adjustment_type"`

	CoolDownMinutes int32 `json:"cool_down_minutes"`

	ScalingAdjustment int32 `json:"scaling_adjustment"`

	Trigger *Trigger `json:"trigger"`
}

func (o Rule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Rule struct{}"
	}

	return strings.Join([]string{"Rule", string(data)}, " ")
}

type RuleAdjustmentType struct {
	value string
}

type RuleAdjustmentTypeEnum struct {
	SCALE_OUT RuleAdjustmentType
	SCALE_IN  RuleAdjustmentType
}

func GetRuleAdjustmentTypeEnum() RuleAdjustmentTypeEnum {
	return RuleAdjustmentTypeEnum{
		SCALE_OUT: RuleAdjustmentType{
			value: "scale_out",
		},
		SCALE_IN: RuleAdjustmentType{
			value: "scale_in",
		},
	}
}

func (c RuleAdjustmentType) Value() string {
	return c.value
}

func (c RuleAdjustmentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RuleAdjustmentType) UnmarshalJSON(b []byte) error {
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
