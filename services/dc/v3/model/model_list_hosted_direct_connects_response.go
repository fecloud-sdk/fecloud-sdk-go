package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListHostedDirectConnectsResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	HostedConnects *[]HostedDirectConnect `json:"hosted_connects,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListHostedDirectConnectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostedDirectConnectsResponse struct{}"
	}

	return strings.Join([]string{"ListHostedDirectConnectsResponse", string(data)}, " ")
}
