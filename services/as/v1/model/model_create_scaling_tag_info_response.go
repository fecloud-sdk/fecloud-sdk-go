package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateScalingTagInfoResponse Response Object
type CreateScalingTagInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateScalingTagInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingTagInfoResponse struct{}"
	}

	return strings.Join([]string{"CreateScalingTagInfoResponse", string(data)}, " ")
}
