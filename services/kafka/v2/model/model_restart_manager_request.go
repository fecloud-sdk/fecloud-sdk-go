package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RestartManagerRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o RestartManagerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartManagerRequest struct{}"
	}

	return strings.Join([]string{"RestartManagerRequest", string(data)}, " ")
}
