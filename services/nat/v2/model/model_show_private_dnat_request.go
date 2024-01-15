package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPrivateDnatRequest struct {
	DnatRuleId string `json:"dnat_rule_id"`
}

func (o ShowPrivateDnatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateDnatRequest struct{}"
	}

	return strings.Join([]string{"ShowPrivateDnatRequest", string(data)}, " ")
}
