package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAz2MigrateRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ListAz2MigrateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAz2MigrateRequest struct{}"
	}

	return strings.Join([]string{"ListAz2MigrateRequest", string(data)}, " ")
}
