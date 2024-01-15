package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteDirectConnectRequest struct {
	DirectConnectId string `json:"direct_connect_id"`
}

func (o DeleteDirectConnectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDirectConnectRequest struct{}"
	}

	return strings.Join([]string{"DeleteDirectConnectRequest", string(data)}, " ")
}
