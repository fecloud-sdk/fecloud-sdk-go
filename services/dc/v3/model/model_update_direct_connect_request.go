package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateDirectConnectRequest struct {
	DirectConnectId string `json:"direct_connect_id"`

	Body *UpdateDirectConnectRequestBody `json:"body,omitempty"`
}

func (o UpdateDirectConnectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDirectConnectRequest struct{}"
	}

	return strings.Join([]string{"UpdateDirectConnectRequest", string(data)}, " ")
}
