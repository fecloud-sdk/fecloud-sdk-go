package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SwitchSslRequestBody struct {
	SslOption string `json:"ssl_option"`
}

func (o SwitchSslRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSslRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchSslRequestBody", string(data)}, " ")
}
