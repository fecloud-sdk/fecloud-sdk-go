package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdatePublicZoneInfo struct {
	Description *string `json:"description,omitempty"`

	Email *string `json:"email,omitempty"`

	Ttl *int32 `json:"ttl,omitempty"`
}

func (o UpdatePublicZoneInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicZoneInfo struct{}"
	}

	return strings.Join([]string{"UpdatePublicZoneInfo", string(data)}, " ")
}
