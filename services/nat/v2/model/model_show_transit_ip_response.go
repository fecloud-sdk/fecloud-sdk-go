package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowTransitIpResponse struct {
	TransitIp *TransitIp `json:"transit_ip,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTransitIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTransitIpResponse struct{}"
	}

	return strings.Join([]string{"ShowTransitIpResponse", string(data)}, " ")
}
