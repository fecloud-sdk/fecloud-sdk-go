package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListTransitIpsResponse struct {
	TransitIps *[]TransitIp `json:"transit_ips,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTransitIpsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransitIpsResponse struct{}"
	}

	return strings.Join([]string{"ListTransitIpsResponse", string(data)}, " ")
}
