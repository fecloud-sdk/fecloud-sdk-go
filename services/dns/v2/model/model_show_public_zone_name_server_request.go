package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPublicZoneNameServerRequest struct {
	ZoneId string `json:"zone_id"`
}

func (o ShowPublicZoneNameServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicZoneNameServerRequest struct{}"
	}

	return strings.Join([]string{"ShowPublicZoneNameServerRequest", string(data)}, " ")
}
