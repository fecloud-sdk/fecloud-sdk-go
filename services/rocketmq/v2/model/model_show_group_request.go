package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowGroupRequest struct {
	InstanceId string `json:"instance_id"`

	Group string `json:"group"`
}

func (o ShowGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowGroupRequest", string(data)}, " ")
}
