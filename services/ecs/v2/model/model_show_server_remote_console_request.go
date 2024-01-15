package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowServerRemoteConsoleRequest struct {
	ServerId string `json:"server_id"`

	Body *ShowServerRemoteConsoleRequestBody `json:"body,omitempty"`
}

func (o ShowServerRemoteConsoleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerRemoteConsoleRequest struct{}"
	}

	return strings.Join([]string{"ShowServerRemoteConsoleRequest", string(data)}, " ")
}
