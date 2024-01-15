package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type FileStatusV2 struct {
	PathSuffix *string `json:"path_suffix,omitempty"`

	Owner *string `json:"owner,omitempty"`

	Group *string `json:"group,omitempty"`

	Permission *string `json:"permission,omitempty"`

	Replication *int32 `json:"replication,omitempty"`

	BlockSize *int32 `json:"block_size,omitempty"`

	Length *int32 `json:"length,omitempty"`

	Type *string `json:"type,omitempty"`

	ChildrenNum *int32 `json:"children_num,omitempty"`

	AccessTime *int64 `json:"access_time,omitempty"`

	ModificationTime *int64 `json:"modification_time,omitempty"`
}

func (o FileStatusV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileStatusV2 struct{}"
	}

	return strings.Join([]string{"FileStatusV2", string(data)}, " ")
}
