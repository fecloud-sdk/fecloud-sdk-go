package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NovaAttachInterfaceRequest struct {
	ServerId string `json:"server_id"`

	Body *NovaAttachInterfaceRequestBody `json:"body,omitempty"`
}

func (o NovaAttachInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaAttachInterfaceRequest struct{}"
	}

	return strings.Join([]string{"NovaAttachInterfaceRequest", string(data)}, " ")
}
