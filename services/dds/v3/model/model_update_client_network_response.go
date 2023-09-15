package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdateClientNetworkResponse Response Object
type UpdateClientNetworkResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateClientNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClientNetworkResponse struct{}"
	}

	return strings.Join([]string{"UpdateClientNetworkResponse", string(data)}, " ")
}
