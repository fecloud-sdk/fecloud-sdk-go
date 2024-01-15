package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteListenerForceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteListenerForceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteListenerForceResponse struct{}"
	}

	return strings.Join([]string{"DeleteListenerForceResponse", string(data)}, " ")
}
