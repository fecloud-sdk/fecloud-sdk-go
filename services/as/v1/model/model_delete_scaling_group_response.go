package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteScalingGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteScalingGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScalingGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteScalingGroupResponse", string(data)}, " ")
}
