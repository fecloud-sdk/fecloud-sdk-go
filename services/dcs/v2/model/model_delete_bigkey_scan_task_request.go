package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteBigkeyScanTaskRequest struct {
	InstanceId string `json:"instance_id"`

	BigkeyId string `json:"bigkey_id"`
}

func (o DeleteBigkeyScanTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBigkeyScanTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteBigkeyScanTaskRequest", string(data)}, " ")
}
