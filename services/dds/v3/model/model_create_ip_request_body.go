package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateIpRequestBody struct {
	Type string `json:"type"`

	TargetId *string `json:"target_id,omitempty"`

	Password string `json:"password"`
}

func (o CreateIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpRequestBody struct{}"
	}

	return strings.Join([]string{"CreateIpRequestBody", string(data)}, " ")
}
