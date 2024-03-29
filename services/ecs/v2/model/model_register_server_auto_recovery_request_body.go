package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RegisterServerAutoRecoveryRequestBody struct {
	SupportAutoRecovery string `json:"support_auto_recovery"`
}

func (o RegisterServerAutoRecoveryRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterServerAutoRecoveryRequestBody struct{}"
	}

	return strings.Join([]string{"RegisterServerAutoRecoveryRequestBody", string(data)}, " ")
}
