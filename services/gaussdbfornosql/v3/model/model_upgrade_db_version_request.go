package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpgradeDbVersionRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o UpgradeDbVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeDbVersionRequest struct{}"
	}

	return strings.Join([]string{"UpgradeDbVersionRequest", string(data)}, " ")
}
