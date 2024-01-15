package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowCoordinatorsRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ShowCoordinatorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCoordinatorsRequest struct{}"
	}

	return strings.Join([]string{"ShowCoordinatorsRequest", string(data)}, " ")
}
