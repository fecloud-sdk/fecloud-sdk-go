package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateHostedDirectConnect struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Bandwidth *int32 `json:"bandwidth,omitempty"`

	PeerLocation *string `json:"peer_location,omitempty"`
}

func (o UpdateHostedDirectConnect) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostedDirectConnect struct{}"
	}

	return strings.Join([]string{"UpdateHostedDirectConnect", string(data)}, " ")
}
