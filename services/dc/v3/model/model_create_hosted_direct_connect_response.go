package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateHostedDirectConnectResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	HostedConnect  *HostedDirectConnect `json:"hosted_connect,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o CreateHostedDirectConnectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostedDirectConnectResponse struct{}"
	}

	return strings.Join([]string{"CreateHostedDirectConnectResponse", string(data)}, " ")
}
