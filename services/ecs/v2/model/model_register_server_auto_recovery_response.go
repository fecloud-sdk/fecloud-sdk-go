package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RegisterServerAutoRecoveryResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RegisterServerAutoRecoveryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterServerAutoRecoveryResponse struct{}"
	}

	return strings.Join([]string{"RegisterServerAutoRecoveryResponse", string(data)}, " ")
}
