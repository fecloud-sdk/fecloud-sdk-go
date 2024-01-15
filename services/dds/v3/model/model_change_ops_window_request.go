package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeOpsWindowRequest struct {
	InstanceId string `json:"instance_id"`

	Body *OpsWindowRequestBody `json:"body,omitempty"`
}

func (o ChangeOpsWindowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeOpsWindowRequest struct{}"
	}

	return strings.Join([]string{"ChangeOpsWindowRequest", string(data)}, " ")
}
