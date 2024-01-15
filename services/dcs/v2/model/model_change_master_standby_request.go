package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeMasterStandbyRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ChangeMasterStandbyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeMasterStandbyRequest struct{}"
	}

	return strings.Join([]string{"ChangeMasterStandbyRequest", string(data)}, " ")
}
