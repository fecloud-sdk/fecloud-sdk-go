package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPrivateZoneNameServerRequest struct {
	ZoneId string `json:"zone_id"`
}

func (o ShowPrivateZoneNameServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateZoneNameServerRequest struct{}"
	}

	return strings.Join([]string{"ShowPrivateZoneNameServerRequest", string(data)}, " ")
}
