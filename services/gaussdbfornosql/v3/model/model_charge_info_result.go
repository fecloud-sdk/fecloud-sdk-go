package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChargeInfoResult struct {
	ChargeMode *string `json:"charge_mode,omitempty"`

	PeriodType *string `json:"period_type,omitempty"`

	PeriodNum *string `json:"period_num,omitempty"`

	IsAutoRenew *string `json:"is_auto_renew,omitempty"`

	IsAutoPay *string `json:"is_auto_pay,omitempty"`
}

func (o ChargeInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChargeInfoResult struct{}"
	}

	return strings.Join([]string{"ChargeInfoResult", string(data)}, " ")
}
