package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListDirectConnectsResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	DirectConnects *[]DirectConnect `json:"direct_connects,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListDirectConnectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDirectConnectsResponse struct{}"
	}

	return strings.Join([]string{"ListDirectConnectsResponse", string(data)}, " ")
}
