package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdatePrivateSnatOption struct {
	TransitIpIds *[]string `json:"transit_ip_ids,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o UpdatePrivateSnatOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateSnatOption struct{}"
	}

	return strings.Join([]string{"UpdatePrivateSnatOption", string(data)}, " ")
}
