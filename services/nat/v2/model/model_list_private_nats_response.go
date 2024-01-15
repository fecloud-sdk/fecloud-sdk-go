package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListPrivateNatsResponse struct {
	Gateways *[]PrivateNat `json:"gateways,omitempty"`

	RequestId *string `json:"request_id,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPrivateNatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateNatsResponse struct{}"
	}

	return strings.Join([]string{"ListPrivateNatsResponse", string(data)}, " ")
}
