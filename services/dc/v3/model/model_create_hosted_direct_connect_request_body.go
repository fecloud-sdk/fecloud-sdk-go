package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateHostedDirectConnectRequestBody struct {
	HostedConnect *CreateHostedDirectConnect `json:"hosted_connect"`
}

func (o CreateHostedDirectConnectRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostedDirectConnectRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHostedDirectConnectRequestBody", string(data)}, " ")
}
