package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteScalingInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteScalingInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScalingInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteScalingInstanceResponse", string(data)}, " ")
}
