package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateEndpointConnectionsDescResponse struct {
	Connections    *[]ConnectionEndpoints `json:"connections,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o UpdateEndpointConnectionsDescResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointConnectionsDescResponse struct{}"
	}

	return strings.Join([]string{"UpdateEndpointConnectionsDescResponse", string(data)}, " ")
}
