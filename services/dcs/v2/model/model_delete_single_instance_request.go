package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteSingleInstanceRequest Request Object
type DeleteSingleInstanceRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o DeleteSingleInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSingleInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteSingleInstanceRequest", string(data)}, " ")
}
