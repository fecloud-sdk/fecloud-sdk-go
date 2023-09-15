package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ResumeScalingGroupResponse Response Object
type ResumeScalingGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResumeScalingGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumeScalingGroupResponse struct{}"
	}

	return strings.Join([]string{"ResumeScalingGroupResponse", string(data)}, " ")
}
