package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateScalingConfigRequest struct {
	Body *CreateScalingConfigOption `json:"body,omitempty"`
}

func (o CreateScalingConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingConfigRequest struct{}"
	}

	return strings.Join([]string{"CreateScalingConfigRequest", string(data)}, " ")
}
