package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateHostedDirectConnectRequest struct {
	Body *CreateHostedDirectConnectRequestBody `json:"body,omitempty"`
}

func (o CreateHostedDirectConnectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostedDirectConnectRequest struct{}"
	}

	return strings.Join([]string{"CreateHostedDirectConnectRequest", string(data)}, " ")
}
