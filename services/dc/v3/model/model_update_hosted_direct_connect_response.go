package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateHostedDirectConnectResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	HostedConnect  *HostedDirectConnect `json:"hosted_connect,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o UpdateHostedDirectConnectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostedDirectConnectResponse struct{}"
	}

	return strings.Join([]string{"UpdateHostedDirectConnectResponse", string(data)}, " ")
}
