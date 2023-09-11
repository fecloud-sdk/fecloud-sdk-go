package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowPrivateSnatRequest Request Object
type ShowPrivateSnatRequest struct {

	// SNAT规则的ID。
	SnatRuleId string `json:"snat_rule_id"`
}

func (o ShowPrivateSnatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateSnatRequest struct{}"
	}

	return strings.Join([]string{"ShowPrivateSnatRequest", string(data)}, " ")
}
