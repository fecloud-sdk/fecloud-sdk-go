package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChargeInfoOption struct {
	ChargeMode string `json:"charge_mode"`

	PeriodType *string `json:"period_type,omitempty"`

	PeriodNum *string `json:"period_num,omitempty"`

	IsAutoRenew *string `json:"is_auto_renew,omitempty"`

	IsAutoPay *string `json:"is_auto_pay,omitempty"`
}

func (o ChargeInfoOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChargeInfoOption struct{}"
	}

	return strings.Join([]string{"ChargeInfoOption", string(data)}, " ")
}
