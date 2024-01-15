package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchProtectScalingInstancesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchProtectScalingInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchProtectScalingInstancesResponse struct{}"
	}

	return strings.Join([]string{"BatchProtectScalingInstancesResponse", string(data)}, " ")
}
