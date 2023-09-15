package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ChangeMasterStandbyRequest Request Object
type ChangeMasterStandbyRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ChangeMasterStandbyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeMasterStandbyRequest struct{}"
	}

	return strings.Join([]string{"ChangeMasterStandbyRequest", string(data)}, " ")
}
