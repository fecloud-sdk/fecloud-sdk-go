package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowResourcesDetailResponseBody struct {
	Type string `json:"type"`

	Mode string `json:"mode"`

	Quota int32 `json:"quota"`

	Used int32 `json:"used"`
}

func (o ShowResourcesDetailResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourcesDetailResponseBody struct{}"
	}

	return strings.Join([]string{"ShowResourcesDetailResponseBody", string(data)}, " ")
}
