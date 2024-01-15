package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateHostedDirectConnectRequest struct {
	HostedConnectId string `json:"hosted_connect_id"`

	Body *UpdateHostedDirectConnectRequestBody `json:"body,omitempty"`
}

func (o UpdateHostedDirectConnectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostedDirectConnectRequest struct{}"
	}

	return strings.Join([]string{"UpdateHostedDirectConnectRequest", string(data)}, " ")
}
