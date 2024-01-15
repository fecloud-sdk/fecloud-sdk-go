package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type BssParam struct {
	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`

	ChargingMode *BssParamChargingMode `json:"charging_mode,omitempty"`

	IsAutoPay *bool `json:"is_auto_pay,omitempty"`

	PeriodType *BssParamPeriodType `json:"period_type,omitempty"`

	PeriodNum *int32 `json:"period_num,omitempty"`

	OrderId *string `json:"order_id,omitempty"`
}

func (o BssParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BssParam struct{}"
	}

	return strings.Join([]string{"BssParam", string(data)}, " ")
}

type BssParamChargingMode struct {
	value string
}

type BssParamChargingModeEnum struct {
	PRE_PAID  BssParamChargingMode
	POST_PAID BssParamChargingMode
}

func GetBssParamChargingModeEnum() BssParamChargingModeEnum {
	return BssParamChargingModeEnum{
		PRE_PAID: BssParamChargingMode{
			value: "prePaid",
		},
		POST_PAID: BssParamChargingMode{
			value: "postPaid",
		},
	}
}

func (c BssParamChargingMode) Value() string {
	return c.value
}

func (c BssParamChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BssParamChargingMode) UnmarshalJSON(b []byte) error {
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

type BssParamPeriodType struct {
	value string
}

type BssParamPeriodTypeEnum struct {
	MONTH BssParamPeriodType
	YEAR  BssParamPeriodType
}

func GetBssParamPeriodTypeEnum() BssParamPeriodTypeEnum {
	return BssParamPeriodTypeEnum{
		MONTH: BssParamPeriodType{
			value: "month",
		},
		YEAR: BssParamPeriodType{
			value: "year",
		},
	}
}

func (c BssParamPeriodType) Value() string {
	return c.value
}

func (c BssParamPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BssParamPeriodType) UnmarshalJSON(b []byte) error {
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
