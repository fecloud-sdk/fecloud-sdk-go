package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdatePrivateZoneInfoReq struct {
	Description *string `json:"description,omitempty"`

	Email *string `json:"email,omitempty"`

	Ttl *int32 `json:"ttl,omitempty"`
}

func (o UpdatePrivateZoneInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateZoneInfoReq struct{}"
	}

	return strings.Join([]string{"UpdatePrivateZoneInfoReq", string(data)}, " ")
}
