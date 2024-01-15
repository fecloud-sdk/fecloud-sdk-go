package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DiskVolumes struct {
	EntityId string `json:"entity_id"`

	EntityName string `json:"entity_name"`

	GroupType string `json:"group_type"`

	Used string `json:"used"`

	Size string `json:"size"`
}

func (o DiskVolumes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskVolumes struct{}"
	}

	return strings.Join([]string{"DiskVolumes", string(data)}, " ")
}
