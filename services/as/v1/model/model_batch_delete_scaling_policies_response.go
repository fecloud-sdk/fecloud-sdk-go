package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteScalingPoliciesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteScalingPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteScalingPoliciesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteScalingPoliciesResponse", string(data)}, " ")
}
