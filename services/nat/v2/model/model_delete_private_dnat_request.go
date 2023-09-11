package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeletePrivateDnatRequest Request Object
type DeletePrivateDnatRequest struct {

	// DNAT规则的ID。
	DnatRuleId string `json:"dnat_rule_id"`
}

func (o DeletePrivateDnatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateDnatRequest struct{}"
	}

	return strings.Join([]string{"DeletePrivateDnatRequest", string(data)}, " ")
}
