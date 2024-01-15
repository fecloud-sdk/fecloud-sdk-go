package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResetManagerPasswordRequest struct {
	InstanceId string `json:"instance_id"`

	Body *ResetManagerPasswordReq `json:"body,omitempty"`
}

func (o ResetManagerPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetManagerPasswordRequest struct{}"
	}

	return strings.Join([]string{"ResetManagerPasswordRequest", string(data)}, " ")
}
