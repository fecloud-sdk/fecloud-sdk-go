package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EntityInfo struct {
	EntityId string `json:"entity_id"`

	EntityName string `json:"entity_name"`
}

func (o EntityInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EntityInfo struct{}"
	}

	return strings.Join([]string{"EntityInfo", string(data)}, " ")
}
