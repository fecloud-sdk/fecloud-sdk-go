package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdatePrivateZoneRequest struct {
	ZoneId string `json:"zone_id"`

	Body *UpdatePrivateZoneInfoReq `json:"body,omitempty"`
}

func (o UpdatePrivateZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateZoneRequest struct{}"
	}

	return strings.Join([]string{"UpdatePrivateZoneRequest", string(data)}, " ")
}
