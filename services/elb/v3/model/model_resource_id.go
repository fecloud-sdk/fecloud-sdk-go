package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ResourceId 资源ID
type ResourceId struct {

	// 资源ID
	Id string `json:"id"`
}

func (o ResourceId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceId struct{}"
	}

	return strings.Join([]string{"ResourceId", string(data)}, " ")
}
