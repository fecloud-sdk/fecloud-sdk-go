package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowCesHierarchyRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ShowCesHierarchyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCesHierarchyRequest struct{}"
	}

	return strings.Join([]string{"ShowCesHierarchyRequest", string(data)}, " ")
}
