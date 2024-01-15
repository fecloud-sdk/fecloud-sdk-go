package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type MrsResource struct {
	ResourceId *string `json:"resource_id,omitempty"`

	ResourceDetail *string `json:"resource_detail,omitempty"`

	Tags *[]TagPlain `json:"tags,omitempty"`

	ResourceName *string `json:"resource_name,omitempty"`
}

func (o MrsResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MrsResource struct{}"
	}

	return strings.Join([]string{"MrsResource", string(data)}, " ")
}
