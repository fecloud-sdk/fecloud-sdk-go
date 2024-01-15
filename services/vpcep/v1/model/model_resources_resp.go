package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResourcesResp struct {
	Resources *[]Quotas `json:"resources,omitempty"`
}

func (o ResourcesResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesResp struct{}"
	}

	return strings.Join([]string{"ResourcesResp", string(data)}, " ")
}
