package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteHostedDirectConnectRequest struct {
	HostedConnectId string `json:"hosted_connect_id"`
}

func (o DeleteHostedDirectConnectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHostedDirectConnectRequest struct{}"
	}

	return strings.Join([]string{"DeleteHostedDirectConnectRequest", string(data)}, " ")
}
