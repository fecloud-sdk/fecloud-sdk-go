package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowHostedDirectConnectResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	HostedConnect  *HostedDirectConnect `json:"hosted_connect,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowHostedDirectConnectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHostedDirectConnectResponse struct{}"
	}

	return strings.Join([]string{"ShowHostedDirectConnectResponse", string(data)}, " ")
}
