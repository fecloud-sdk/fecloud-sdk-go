package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type VersionItem struct {
	Id *string `json:"id,omitempty"`

	Status *string `json:"status,omitempty"`

	Links *[]LinksItem `json:"links,omitempty"`

	Updated *string `json:"updated,omitempty"`

	Version *string `json:"version,omitempty"`

	MinVersion *string `json:"min_version,omitempty"`
}

func (o VersionItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionItem struct{}"
	}

	return strings.Join([]string{"VersionItem", string(data)}, " ")
}
