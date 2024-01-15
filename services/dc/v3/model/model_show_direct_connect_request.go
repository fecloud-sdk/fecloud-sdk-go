package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowDirectConnectRequest struct {
	DirectConnectId string `json:"direct_connect_id"`

	Fields *[]string `json:"fields,omitempty"`
}

func (o ShowDirectConnectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDirectConnectRequest struct{}"
	}

	return strings.Join([]string{"ShowDirectConnectRequest", string(data)}, " ")
}
