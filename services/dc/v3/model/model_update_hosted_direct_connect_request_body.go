package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateHostedDirectConnectRequestBody struct {
	HostedConnect *UpdateHostedDirectConnect `json:"hosted_connect,omitempty"`
}

func (o UpdateHostedDirectConnectRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostedDirectConnectRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHostedDirectConnectRequestBody", string(data)}, " ")
}
