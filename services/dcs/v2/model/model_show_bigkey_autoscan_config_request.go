package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowBigkeyAutoscanConfigRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ShowBigkeyAutoscanConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBigkeyAutoscanConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowBigkeyAutoscanConfigRequest", string(data)}, " ")
}
