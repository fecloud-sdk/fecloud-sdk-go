package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowGroupsRequest struct {
	InstanceId string `json:"instance_id"`

	Group string `json:"group"`
}

func (o ShowGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupsRequest struct{}"
	}

	return strings.Join([]string{"ShowGroupsRequest", string(data)}, " ")
}
