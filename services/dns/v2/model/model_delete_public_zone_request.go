package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeletePublicZoneRequest struct {
	ZoneId string `json:"zone_id"`
}

func (o DeletePublicZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePublicZoneRequest struct{}"
	}

	return strings.Join([]string{"DeletePublicZoneRequest", string(data)}, " ")
}
