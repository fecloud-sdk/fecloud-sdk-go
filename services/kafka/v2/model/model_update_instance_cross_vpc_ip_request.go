package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceCrossVpcIpRequest struct {
	InstanceId string `json:"instance_id"`

	Body *UpdateInstanceCrossVpcIpReq `json:"body,omitempty"`
}

func (o UpdateInstanceCrossVpcIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceCrossVpcIpRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceCrossVpcIpRequest", string(data)}, " ")
}
