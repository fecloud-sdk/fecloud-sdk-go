package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdatePrivateSnatRequest struct {
	SnatRuleId string `json:"snat_rule_id"`

	Body *UpdatePrivateSnatOptionBody `json:"body,omitempty"`
}

func (o UpdatePrivateSnatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateSnatRequest struct{}"
	}

	return strings.Join([]string{"UpdatePrivateSnatRequest", string(data)}, " ")
}
