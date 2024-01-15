package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateEpConnections struct {
	Connections []ConnectionsDesc `json:"connections"`
}

func (o UpdateEpConnections) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEpConnections struct{}"
	}

	return strings.Join([]string{"UpdateEpConnections", string(data)}, " ")
}
