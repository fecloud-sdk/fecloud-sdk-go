package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpgradeDatabaseVersionRequestBody struct {
	UpgradeMode *string `json:"upgrade_mode,omitempty"`
}

func (o UpgradeDatabaseVersionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeDatabaseVersionRequestBody struct{}"
	}

	return strings.Join([]string{"UpgradeDatabaseVersionRequestBody", string(data)}, " ")
}
