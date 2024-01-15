package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChargeInfoResult struct {
	ChargeMode *string `json:"charge_mode,omitempty"`

	PeriodType *string `json:"period_type,omitempty"`

	PeriodNum *int32 `json:"period_num,omitempty"`

	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`

	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o ChargeInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChargeInfoResult struct{}"
	}

	return strings.Join([]string{"ChargeInfoResult", string(data)}, " ")
}
