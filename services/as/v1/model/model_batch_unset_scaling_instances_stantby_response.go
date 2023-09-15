package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchUnsetScalingInstancesStantbyResponse Response Object
type BatchUnsetScalingInstancesStantbyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchUnsetScalingInstancesStantbyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUnsetScalingInstancesStantbyResponse struct{}"
	}

	return strings.Join([]string{"BatchUnsetScalingInstancesStantbyResponse", string(data)}, " ")
}
