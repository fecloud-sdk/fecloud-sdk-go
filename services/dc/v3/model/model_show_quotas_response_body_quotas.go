package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowQuotasResponseBodyQuotas struct {
	Resources *[]Info `json:"resources,omitempty"`
}

func (o ShowQuotasResponseBodyQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotasResponseBodyQuotas struct{}"
	}

	return strings.Join([]string{"ShowQuotasResponseBodyQuotas", string(data)}, " ")
}
