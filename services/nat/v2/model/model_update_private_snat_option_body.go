package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdatePrivateSnatOptionBody struct {
	SnatRule *UpdatePrivateSnatOption `json:"snat_rule"`
}

func (o UpdatePrivateSnatOptionBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateSnatOptionBody struct{}"
	}

	return strings.Join([]string{"UpdatePrivateSnatOptionBody", string(data)}, " ")
}
