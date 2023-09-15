package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpgradeDatabaseVersionRequest Request Object
type UpgradeDatabaseVersionRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *UpgradeDatabaseVersionRequestBody `json:"body,omitempty"`
}

func (o UpgradeDatabaseVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeDatabaseVersionRequest struct{}"
	}

	return strings.Join([]string{"UpgradeDatabaseVersionRequest", string(data)}, " ")
}
