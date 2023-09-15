package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchDeleteScalingConfigsResponse Response Object
type BatchDeleteScalingConfigsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteScalingConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteScalingConfigsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteScalingConfigsResponse", string(data)}, " ")
}
