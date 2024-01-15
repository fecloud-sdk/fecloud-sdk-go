package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CheckPasswordRequest struct {
	InstanceId string `json:"instance_id"`

	Body *CheckPasswordRequestBody `json:"body,omitempty"`
}

func (o CheckPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPasswordRequest struct{}"
	}

	return strings.Join([]string{"CheckPasswordRequest", string(data)}, " ")
}
