package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SwitchHostsProtectStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchHostsProtectStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchHostsProtectStatusResponse struct{}"
	}

	return strings.Join([]string{"SwitchHostsProtectStatusResponse", string(data)}, " ")
}
