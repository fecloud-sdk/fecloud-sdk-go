package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowDirectConnectResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	DirectConnect  *DirectConnect `json:"direct_connect,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowDirectConnectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDirectConnectResponse struct{}"
	}

	return strings.Join([]string{"ShowDirectConnectResponse", string(data)}, " ")
}
