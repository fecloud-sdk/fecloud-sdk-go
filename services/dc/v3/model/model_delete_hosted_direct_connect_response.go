package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteHostedDirectConnectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHostedDirectConnectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHostedDirectConnectResponse struct{}"
	}

	return strings.Join([]string{"DeleteHostedDirectConnectResponse", string(data)}, " ")
}
