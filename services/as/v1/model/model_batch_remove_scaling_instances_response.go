package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchRemoveScalingInstancesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchRemoveScalingInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemoveScalingInstancesResponse struct{}"
	}

	return strings.Join([]string{"BatchRemoveScalingInstancesResponse", string(data)}, " ")
}
