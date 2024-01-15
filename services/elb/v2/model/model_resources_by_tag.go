package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResourcesByTag struct {
	ResourceId string `json:"resource_id"`

	ResourceName string `json:"resource_name"`

	ResourceDetail string `json:"resource_detail"`

	Tags []ResourceTag `json:"tags"`

	SuperResourceId *string `json:"super_resource_id,omitempty"`
}

func (o ResourcesByTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesByTag struct{}"
	}

	return strings.Join([]string{"ResourcesByTag", string(data)}, " ")
}
