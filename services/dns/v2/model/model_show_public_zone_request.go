package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPublicZoneRequest struct {
	ZoneId string `json:"zone_id"`
}

func (o ShowPublicZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicZoneRequest struct{}"
	}

	return strings.Join([]string{"ShowPublicZoneRequest", string(data)}, " ")
}
