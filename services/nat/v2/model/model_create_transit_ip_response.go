package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateTransitIpResponse struct {
	TransitIp *TransitIp `json:"transit_ip,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTransitIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransitIpResponse struct{}"
	}

	return strings.Join([]string{"CreateTransitIpResponse", string(data)}, " ")
}
