package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPrivateZoneRequest struct {
	ZoneId string `json:"zone_id"`
}

func (o ShowPrivateZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateZoneRequest struct{}"
	}

	return strings.Join([]string{"ShowPrivateZoneRequest", string(data)}, " ")
}
