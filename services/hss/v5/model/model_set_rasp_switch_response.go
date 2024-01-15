package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SetRaspSwitchResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetRaspSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRaspSwitchResponse struct{}"
	}

	return strings.Join([]string{"SetRaspSwitchResponse", string(data)}, " ")
}
