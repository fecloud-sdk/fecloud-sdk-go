package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteDirectConnectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDirectConnectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDirectConnectResponse struct{}"
	}

	return strings.Join([]string{"DeleteDirectConnectResponse", string(data)}, " ")
}
