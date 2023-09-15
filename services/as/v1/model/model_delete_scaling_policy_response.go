package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteScalingPolicyResponse Response Object
type DeleteScalingPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteScalingPolicyResponse", string(data)}, " ")
}
