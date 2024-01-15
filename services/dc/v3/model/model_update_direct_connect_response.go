package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateDirectConnectResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	DirectConnect  *DirectConnect `json:"direct_connect,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateDirectConnectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDirectConnectResponse struct{}"
	}

	return strings.Join([]string{"UpdateDirectConnectResponse", string(data)}, " ")
}
