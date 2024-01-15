package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowConsumerConnectionsResponse struct {
	GroupName *string `json:"group_name,omitempty"`

	Online *bool `json:"online,omitempty"`

	SubscriptionConsistency *bool `json:"subscription_consistency,omitempty"`

	Total *int32 `json:"total,omitempty"`

	NextOffset *int32 `json:"next_offset,omitempty"`

	PreviousOffset *int32 `json:"previous_offset,omitempty"`

	Clients        *[]ClientData `json:"clients,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowConsumerConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsumerConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ShowConsumerConnectionsResponse", string(data)}, " ")
}
