package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type MigrateAzRequest struct {
	InstanceId string `json:"instance_id"`

	Body *MigrateAzRequestBody `json:"body,omitempty"`
}

func (o MigrateAzRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateAzRequest struct{}"
	}

	return strings.Join([]string{"MigrateAzRequest", string(data)}, " ")
}
