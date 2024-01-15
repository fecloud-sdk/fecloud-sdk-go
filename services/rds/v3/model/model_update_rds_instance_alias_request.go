package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateRdsInstanceAliasRequest struct {
	Alias *string `json:"alias,omitempty"`
}

func (o UpdateRdsInstanceAliasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRdsInstanceAliasRequest struct{}"
	}

	return strings.Join([]string{"UpdateRdsInstanceAliasRequest", string(data)}, " ")
}
