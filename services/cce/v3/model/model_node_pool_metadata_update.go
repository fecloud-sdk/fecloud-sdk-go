package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NodePoolMetadataUpdate struct {
	Name string `json:"name"`
}

func (o NodePoolMetadataUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePoolMetadataUpdate struct{}"
	}

	return strings.Join([]string{"NodePoolMetadataUpdate", string(data)}, " ")
}
