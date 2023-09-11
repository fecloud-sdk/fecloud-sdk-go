package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type JobMetadata struct {
	Uid *string `json:"uid,omitempty"`

	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	UpdateTimestamp *string `json:"updateTimestamp,omitempty"`
}

func (o JobMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobMetadata struct{}"
	}

	return strings.Join([]string{"JobMetadata", string(data)}, " ")
}
