package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdatePublicZoneRequest struct {
	ZoneId string `json:"zone_id"`

	Body *UpdatePublicZoneInfo `json:"body,omitempty"`
}

func (o UpdatePublicZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicZoneRequest struct{}"
	}

	return strings.Join([]string{"UpdatePublicZoneRequest", string(data)}, " ")
}
