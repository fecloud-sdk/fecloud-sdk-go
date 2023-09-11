package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// NovaCreateServersResponse Response Object
type NovaCreateServersResponse struct {
	Server         *NovaCreateServersResult `json:"server,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o NovaCreateServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaCreateServersResponse struct{}"
	}

	return strings.Join([]string{"NovaCreateServersResponse", string(data)}, " ")
}
