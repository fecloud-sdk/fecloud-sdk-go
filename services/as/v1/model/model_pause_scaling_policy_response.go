package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PauseScalingPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o PauseScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"PauseScalingPolicyResponse", string(data)}, " ")
}
