package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeletePrivateZoneRequest struct {
	ZoneId string `json:"zone_id"`
}

func (o DeletePrivateZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateZoneRequest struct{}"
	}

	return strings.Join([]string{"DeletePrivateZoneRequest", string(data)}, " ")
}
