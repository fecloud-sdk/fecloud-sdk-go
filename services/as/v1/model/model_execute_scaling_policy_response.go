package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ExecuteScalingPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"ExecuteScalingPolicyResponse", string(data)}, " ")
}
