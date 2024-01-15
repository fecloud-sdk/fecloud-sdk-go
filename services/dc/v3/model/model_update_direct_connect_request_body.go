package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateDirectConnectRequestBody struct {
	DirectConnect *UpdateDirectConnect `json:"direct_connect,omitempty"`
}

func (o UpdateDirectConnectRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDirectConnectRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDirectConnectRequestBody", string(data)}, " ")
}
