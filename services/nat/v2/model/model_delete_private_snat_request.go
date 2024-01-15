package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeletePrivateSnatRequest struct {
	SnatRuleId string `json:"snat_rule_id"`
}

func (o DeletePrivateSnatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateSnatRequest struct{}"
	}

	return strings.Join([]string{"DeletePrivateSnatRequest", string(data)}, " ")
}
